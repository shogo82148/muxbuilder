package main

import "fmt"

func Get(context *GetContext) error {
	fmt.Fprint(context.ResponseWriter, "Your request method is GET")
	return nil
}

func Post(context *PostContext) error {
	fmt.Fprint(context.ResponseWriter, "Your request method is POST")
	return nil
}
