package main
import (
    "fmt"
)
func main(){
    // array and slices
    // array is fixed size
    // slice is dynamic size
    // declare array
    var ages [3]int = [3]int{1,2,3}
    fmt.Println(ages, len(ages))
    names :=[4]string{"kiya", "Tamirat", "you"}
    fmt.Print(names)
    fmt.Println(len(names))
    scores := []int{1,2,3,4,5,6}
    scores[0] = 1000
    scores = append(scores, 43)
    fmt.Println(scores[2])

   
}