package muxbuilder

import (
	"fmt"
	"io"
)

type handlerDefinitionParam struct {
	name string
	expr string
}

type handlerDefinition struct {
	method      string
	url         string
	name        string
	contextName string
	params      []handlerDefinitionParam
}

func newHandlerDefinition(method, url, name, contextName string) *handlerDefinition {
	return &handlerDefinition{
		method:      method,
		url:         url,
		name:        name,
		contextName: contextName,
	}
}

func (hd *handlerDefinition) write(w io.Writer) {
	fmt.Fprintf(w, "mux.%s(", hd.method)
	fmt.Fprintf(w, "\"%s\", ", hd.url)
	fmt.Fprintln(w, "func(argWriter http.ResponseWriter, argRequest *http.Request, argParams denco.Params) {")
	for _, param := range hd.params {
		fmt.Fprintf(w, "param%s := %s\n", param.name, param.expr)
	}
	fmt.Fprintf(w, "%s(&%s{\n", hd.name, hd.contextName)
	fmt.Fprintln(w, `	Context: Context {
ResponseWriter: argWriter,
Request:        argRequest,
},`)
	for _, param := range hd.params {
		fmt.Fprintf(w, "%s: param%s,\n", param.name, param.name)
	}
	fmt.Fprintln(w, "})")
	fmt.Fprintln(w, "}),")
}

func (hd *handlerDefinition) addParam(name, expr string) {
	hd.params = append(hd.params, handlerDefinitionParam{name, expr})
}
