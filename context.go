package muxbuilder

import (
	"fmt"
	"io"
)

type contextDefinitionField struct {
	name     string
	typename string
}

type contextDefinition struct {
	name   string
	fields []contextDefinitionField
}

func newContextDefinition(name string) *contextDefinition {
	return &contextDefinition{
		name:   name,
		fields: []contextDefinitionField{},
	}
}

func (cd *contextDefinition) addField(name, typename string) {
	field := contextDefinitionField{
		name:     name,
		typename: typename,
	}
	cd.fields = append(cd.fields, field)
}

func (cd *contextDefinition) write(w io.Writer) error {
	fmt.Fprintf(w, "type %s struct {\n", cd.name)
	fmt.Fprintln(w, "Context")
	for _, field := range cd.fields {
		fmt.Fprintf(w, "%s %s\n", field.name, field.typename)
	}
	fmt.Fprintln(w, "}")
	return nil
}
