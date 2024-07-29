package main

import (
    "fmt"
    "strings"
    "sort" 
)

func main() {
    nums := []int{4,2,1,3,4,5,54,643}
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
    // nums.sort()
    sort.Ints(nums)
    fmt.Println(nums)
}