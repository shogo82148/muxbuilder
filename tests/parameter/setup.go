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
				URL:  "/foo/:Name/bar",
				Name: "FooBar",
			},
		},
	}
	ioutil.WriteFile("mux.go", []byte(builder.Build()), 0644)
}
