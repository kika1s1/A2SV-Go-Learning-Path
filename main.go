package main

import "fmt"

func main() {
    i := 0
    for i <= 10 {
        if i % 2 == 1 {
            fmt.Println(i)
        }else if i == 4{
            i++
            continue
        } else{
            break
        }
        i++
    }
}
