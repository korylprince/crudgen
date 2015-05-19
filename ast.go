package crudgen

import (
	"fmt"
	"strings"
)

//Field represents a parsed struct field
type Field struct {
	Name       string
	Type       string
	SQLName    string
	Key        bool
	ForeignKey string
}

//Struct represents a parsed go struct with some metadata
type Struct struct {
	Name        string
	SQLName     string
	Fields      []Field
	Package     string
	DBInterface string //go type for db parameter
}

//Keys returns all of the Struct Fields that are keys
func (s *Struct) Keys() []Field {
	var flds []Field
	for _, f := range s.Fields {
		if f.Key {
			flds = append(flds, f)
		}
	}
	return flds
}

//NameList returns fields in the format "f1.Name, f2.Name, f3.Name, ..., fn.Name"
func NameList(fields []Field) string {
	var names []string
	for _, f := range fields {
		names = append(names, f.Name)
	}
	return strings.Join(names, ", ")
}

//SQLNameList returns fields in the format "f1.SQLName, f2.SQLName, f3.SQLName, ..., fn.SQLName"
func SQLNameList(fields []Field) string {
	var names []string
	for _, f := range fields {
		names = append(names, f.SQLName)
	}
	return strings.Join(names, ", ")
}

//FuncParameterList returns fields in the format
//"f1.Name f1.Type, f2.Name f2.Type, f3.Name f3.Type, ..., fn.Name fn.Type"
func FuncParameterList(fields []Field) string {
	var params []string
	for _, f := range fields {
		params = append(params, fmt.Sprintf("%s %s", f.Name, f.Type))
	}
	return strings.Join(params, ", ")
}

//SQLParameterList returns fields in the format "f1.SQLName=?, f2.SQLName=?, f3.SQLName=?, ..., fn.SQLName=?"
func SQLParameterList(fields []Field) string {
	var params []string
	for _, f := range fields {
		params = append(params, fmt.Sprintf("%s=?", f.SQLName))
	}
	return strings.Join(params, "AND ")
}

//ScanList returns fields in the format
//"&(objName.f1.Name), &(objName.f2.Name), &(objName.f3.Name), ..., &(objName.fn.Name)"
func ScanList(objName string, fields []Field) string {
	var params []string
	for _, f := range fields {
		params = append(params, fmt.Sprintf("&(%s.%s)", objName, f.Name))
	}
	return strings.Join(params, ", ")
}
