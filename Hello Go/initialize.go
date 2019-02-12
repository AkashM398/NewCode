package main

import "fmt"

var a, b string = "Hello", "World"

//A var declaration can include initializers, one per variable.
func main() {
	i := 10
	//:= short assignment
	var c, hey = true, "there"
	//initialised with datatype no need dtatype
	fmt.Println(a, b, hey, c)
	fmt.Println(i)
}
