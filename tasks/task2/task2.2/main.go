package main

import (
	"fmt"
	"regexp"
	"strings"
	"bufio"
	"os"
)

func isPalindrome(input string) bool {
    re := regexp.MustCompile(`[^a-zA-Z0-9]`)
    cleanedInput := re.ReplaceAllString(input, "")
    cleanedInput = strings.ToLower(cleanedInput)
    length := len(cleanedInput)
    for i := 0; i < length/2; i++ {
        if cleanedInput[i] != cleanedInput[length-1-i] {
            return false
        }
    }
    return true
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
    fmt.Println(isPalindrome(input)) 
}