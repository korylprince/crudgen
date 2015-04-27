package crudgen

import (
	"bytes"
	"fmt"
	"go/format"
	"text/template"
)

//go:generate go-bindata -pkg crudgen -o tmpl.go tmpl/...

func ExecuteTemplate(tmplName string, s *Struct) ([]byte, error) {
	tmplData, err := Asset(fmt.Sprintf("tmpl/%s.tmpl", tmplName))
	if err != nil {
		return nil, err
	}

	tmpl := template.New(tmplName)
	tmpl.Funcs(template.FuncMap{
		"NameList":          NameList,
		"SQLNameList":       SQLNameList,
		"FuncParameterList": FuncParameterList,
		"SQLParameterList":  SQLParameterList,
		"ScanList":          ScanList,
	})

	_, err = tmpl.Parse(string(tmplData))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, s)

	return buf.Bytes(), err
}

func Generate(tmpls []string, structs []*Struct) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("package %s", structs[0].Package))
	buf.WriteString("\nimport (\n\t\"database/sql\"\n\t\"fmt\"\n)")

	for _, s := range structs {
		for _, tmpl := range tmpls {
			buf.WriteByte('\n')
			b, err := ExecuteTemplate(tmpl, s)
			if err != nil {
				return nil, err
			}
			buf.Write(b)
		}
	}

	return format.Source(buf.Bytes())
}
