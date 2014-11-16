package main

import "fmt"

func Root(context *RootContext) error {
	fmt.Fprint(context.ResponseWriter, "This is Root handler")
	return nil
}

func Foo(context *FooContext) error {
	fmt.Fprint(context.ResponseWriter, "This is Foo handler")
	return nil
}

func Bar(context *BarContext) error {
	fmt.Fprint(context.ResponseWriter, "This is Bar handler")
	return nil
}
