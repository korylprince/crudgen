package crudgen

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"reflect"
	"strings"

	"go/ast"
)

//Parse searches all of the files ending in ".go" in the given directory for the given struct
//Parse returns a Struct if a matching struct is found or an error if one occurred.
func Parse(dir, structName string) (*Struct, error) {
	fset := token.NewFileSet()
	pkgset, err := parser.ParseDir(fset, dir, nil, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	v := &visitor{
		Struct: &Struct{
			Name:   structName,
			Fields: make([]Field, 0),
		},
	}

	for _, pkg := range pkgset {
		for _, f := range pkg.Files {
			ast.Walk(v, f)

			if v.err != nil {
				return nil, v.err
			}

			if v.found {
				v.Struct.Package = pkg.Name
				return v.Struct, nil
			}
		}
	}

	return nil, fmt.Errorf("struct %s not found", structName)
}

func sqlName(givenName, tagName string) string {
	if tagName != "" {
		return tagName
	}
	return strings.ToLower(givenName)
}

type visitor struct {
	found bool
	err   error

	Struct *Struct
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	var (
		decl  *ast.GenDecl
		tSpec *ast.TypeSpec
		ok    bool
	)

	//skip if not a declartion
	if decl, ok = node.(*ast.GenDecl); !ok {
		return v //continue
	}
	for _, s := range decl.Specs {

		//skip if not a type declartion
		if tSpec, ok = s.(*ast.TypeSpec); !ok {
			return v //continue
		}

		//find struct with matching name
		if st, ok := (tSpec.Type).(*ast.StructType); ok && tSpec.Name.String() == v.Struct.Name {

			v.found = true

			//check for empty struct
			if st.Fields == nil || len(st.Fields.List) == 0 {
				v.err = fmt.Errorf("Empty struct %v at %v", v.Struct.Name, st.Pos())
				return nil //stop
			}

			//loop over fields
			for _, f := range st.Fields.List {

				typeBuf := new(bytes.Buffer)
				fset := token.NewFileSet()
				err := printer.Fprint(typeBuf, fset, f.Type)
				if err != nil {
					v.err = err
					return nil //stop
				}

				var key bool
				var fieldName, foreignKey string

				//parse struct tags
				if f.Tag != nil {
					structTag := reflect.StructTag(f.Tag.Value[1:len(f.Tag.Value)])
					if structTag.Get("crud_key") == "true" {
						key = true
					}
					fieldName = structTag.Get("crud_name")
					foreignKey = structTag.Get("crud_foreign_key")

					// set table name if given
					if tableName := structTag.Get("crud_table"); tableName != "" {
						v.Struct.SQLName = tableName
					}
				}

				//add field
				for _, name := range f.Names {
					v.Struct.Fields = append(v.Struct.Fields, Field{
						Name:       name.String(),
						Type:       typeBuf.String(),
						Key:        key,
						SQLName:    sqlName(name.String(), fieldName),
						ForeignKey: foreignKey,
					})
				}

			}

			//check if any key is set
			hasKey := false
			for _, f := range v.Struct.Fields {
				if f.Key {
					hasKey = true
				}
			}

			//if no key is set, set one for ID if it exists are return error if not
			if !hasKey {
				isSet := false
				for i, f := range v.Struct.Fields {
					if strings.ToLower(f.Name) == "id" {
						v.Struct.Fields[i].Key = true
						isSet = true
					}
				}
				if !isSet {
					v.err = fmt.Errorf("No key is set on struct %v at %v", v.Struct.Name, st.Pos())
				}
			}

			//set table name if not set already
			v.Struct.SQLName = sqlName(v.Struct.Name, v.Struct.SQLName)

			return nil //stop
		}
	}
	return v //continue
}
