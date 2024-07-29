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
    // var nameOne string = "kebede"
    // fmt.Print(nameOne)
    // var nameTwo = "kiya"
    // fmt.Print(nameTwo)
    // for future use
    // var nameThree string
    // nameThree = "Sena"
    // fmt.Print(nameThree)
    // fmt.Print(nameOne, nameTwo, nameThree)
    // nameFive := "this is just shorthand "
    // fmt.Print(nameFive)

    // working with numbers 
    var ageOne int8 = 20
    var ageTwo int16 = 32
    ageThree := int32(40)
    fmt.Println(ageOne, ageTwo, ageThree)

    // variation of int types
    var ageFour int64 = 64
    var ageFive uint8 = 8
    ageSix := uint16(16)
    // unsigned integers positive integers only
    fmt.Println(ageFour, ageFive, ageSix)

    // float numbers
    var floatOne float32 = 3.14
    var floatTwo float64 = 6.28
    floatThree := float32(9.87)
    fmt.Println(floatOne, floatTwo, floatThree)
}