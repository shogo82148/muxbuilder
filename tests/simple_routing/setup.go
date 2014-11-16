package main

import (
	"github.com/shogo82148/muxbuilder"
	"io/ioutil"
)

func main() {
	builder := &muxbuilder.MUXBuilder{
		Package: "main",
		Routes: []muxbuilder.Route{
			muxbuilder.Route{
				URL:  "/",
				Name: "Root",
			},
			muxbuilder.Route{
				URL:  "/foo",
				Name: "Foo",
			},
			muxbuilder.Route{
				URL:  "/bar",
				Name: "Bar",
			},
		},
	}
	ioutil.WriteFile("mux.go", []byte(builder.Build()), 0644)
}
