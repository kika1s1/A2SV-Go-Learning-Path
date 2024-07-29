package main
import (
    "fmt"
)

func main(){
    // fmt package
    // this does not allow new line
    fmt.Print("Hello")
    fmt.Print("world")
    // this allows new line
    fmt.Println("Hello")
    fmt.Println("world")
    // formatting string
    // fmt.Printf("Hello %s", "world")
    // %s is for string
    // %d is for integer
    // %f is for float
    // %t is for boolean
    fmt.Printf("Hello %s, your age is %d, and your weight is %.2f kg\n", "world", 30, 75.5)
    fmt.Printf("Is your name %t\n", true)
    fmt.Printf("Is your name %t\n", false)
    // %v is for any type
    fmt.Printf("Is your name %v\n", true)
    fmt.Printf("Is your name %v\n", false)
}