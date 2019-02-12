package main

//Every Go program is made up of Packages
import (
	"fmt"
	"math"
	"math/rand"
	//imported full package
	//imported random from math package if not used shows erroe
)

func main() {
	fmt.Println("Hello, 世界",
		rand.Intn(10), math.Pi)
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(016))
	fmt.Println(add(42, 8))
}

//func functioname (args)return type{}//function
func add(x, y int) int {
	return x + y
}
