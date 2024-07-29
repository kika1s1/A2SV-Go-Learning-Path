package main

import (
    "fmt"
)

func main() {
    // if else
    age :=12
    if age > 18 {
        fmt.Println("You are an adult")
    } else if age ==12 {
        fmt.Println("You are a teenager")
    }else {
        fmt.Println("You are a child")
    }

}