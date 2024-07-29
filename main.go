package main

import (
    "fmt"
    "strings"
)

func main() {
    nums := [3]int{1,2,3}
    greetings := "hello this is tamirat kebede"
    fmt.Println(strings.Contains(greetings, "hello"))
    fmt.Println(strings.ReplaceAll(greetings, "e", "*"))
    fmt.Println(strings.ToUpper(greetings))
    fmt.Println(greetings)
    fmt.Println(strings.Index(greetings, "hello"))
    fmt.Println(strings.Split(greetings, "l"))
    fmt.Println(nums)
    for i:=0; i < len(nums); i++{
        fmt.Println(nums[i])
    }
}