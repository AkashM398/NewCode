package main

import "fmt"
//A function can return any number of results. WOW really
func swap(x, y string) (string, string) {
	return y, x
}
//Naming return values of funnctions
func add(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(swap(b,a))
	fmt.Println(add(17))
	fmt.Println("Really these features of Go are literally Awesome!!!")
}
