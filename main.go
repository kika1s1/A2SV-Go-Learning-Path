package main

import "fmt"
func main() {
	m := make(map[string]int)
	// m := map[string]int{}


	m["Answer"] = 42
    m["this"] = 0
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
    if m["Answer"] == m["this"]{
        fmt.Print("YES")
    }
    
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
    fmt.Println(m)
}
