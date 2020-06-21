package main

import "fmt"

func main() {
	var v interface{}
	fmt.Printf("%#v\n", v) // <nil>

	var p *int
	v = p
	fmt.Printf("%#v\n", v) // (*int)(nil)
}