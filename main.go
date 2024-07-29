package main

import (
    "fmt"
)

func main() {
    // boolean 
    // true or false
    var b bool = true
    fmt.Println(b)
    // condition
    if b {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
    // switch statement
    // switch expression
    // expression is evaluated once
    // then match with case
    // if no case match then default case
    switch 1 + 2 {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("default")
    }
    // fallthrough keyword
    // it allows to fall through to next case
    switch 1 + 2 {
    case 1:
        fmt.Println("one")
        fallthrough
    case 2:
        fmt.Println("two")
    default:
        fmt.Println("default")
    }

}