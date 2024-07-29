package main

import (
    "fmt"
)

func main() {
    // looping in go
    // for loop
    // syntax
    // for initialization; condition; post {
    //     // block of code
    // }
    // initialization
    i := 0
    // condition
    for i < 10 {
        // block of code
        fmt.Println(i)
        // post
        i++
    }
    // loop without initialization
    for j := 0; j < 10; j++ {
        fmt.Println(j)
    }
    // loop without condition
    for {
        fmt.Println("infinite loop")
        break
    }
    // loop with break statement
    for k := 0; k < 10; k++ {
        if k == 5 {
            break
        }
        fmt.Println(k)
    }
    // loop with continue statement
    for l := 0; l < 10; l++ {
        if l % 2 == 0 {
            continue
        }
        fmt.Println(l)
    }
    // loop with label
    outerLoop:
    for m := 0; m < 10; m++ {
        for n := 0; n < 10; n++ {
            if n == 5 {
                break outerLoop
            }
            fmt.Println(m, n)
        }
    }

}