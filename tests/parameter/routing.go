package main

import "fmt"

func FooBar(context *FooBarContext) error {
	fmt.Fprintf(context.ResponseWriter, "My name is %s", context.Name)
	return nil
}
