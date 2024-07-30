package main

import (
    "fmt"
)

func main() {
    information := []struct {
        name string
        age  int
    }{
        {"John", 25},
        {"Jane", 30},
        {"Mike", 35},
    }
    // Add more information to the information variable if needed
    
    information = append(information, struct {
        name string
        age  int
    }{
        name: "Sarah",
        age:  40,
    }, struct {
        name string
        age  int
    }{
        name: "David",
        age:  45,
    })
    fmt.Println("Name\tAge")
    for _, person := range information {
        fmt.Printf("%s\t%d\n", person.name, person.age)
    }
}