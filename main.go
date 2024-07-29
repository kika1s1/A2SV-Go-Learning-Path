package main
import (
    "fmt"
)

// to run go file 
    // - go run main.go
// main function in entry function
// only main function one in one package
// what is fmt ?
// it is just formatting package
// everything in fmt dot should start with Capital letter
// example
        // - fmt.Println("hello world")
        // - fmt.Print("hello world")
func main(){
    // fmt.Print("Initial Go lang")
    // this is just for format string or output code
    // strings are double qoute
    var nameOne string = "kebede"
    // fmt.Print(nameOne)
    var nameTwo = "kiya"
    // fmt.Print(nameTwo)
    // for future use
    var nameThree string
    nameThree = "Sena"
    // fmt.Print(nameThree)
    fmt.Print(nameOne, nameTwo, nameThree)
    nameFive := "this is just shorthand "
    fmt.Print(nameFive)

    

}