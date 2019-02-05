package main

import "fmt"
import "math/rand"
//IMPORT TIME
import "time"

func main() {
    rand.Seed(time.Now().Unix())
    //Specially geneerates random numbers because without this you'll see same series of random numbers generated
    for i := 1; i <= 5; i++ {
        // Generate and print a random integer from 0 to 9
        fmt.Println(rand.Intn(10))
    }
}
