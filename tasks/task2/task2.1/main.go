package main

import (
    "fmt"
    "regexp"
    "strings"
	"bufio"
	"os"
)

func wordFrequencyCount(input string) map[string]int {
    re := regexp.MustCompile(`[^\w\s]`)
    cleanedInput := re.ReplaceAllString(input, "")
    cleanedInput = strings.ToLower(cleanedInput)
    words := strings.Fields(cleanedInput)
    frequency := make(map[string]int)
    for _, word := range words {
        frequency[word] +=1
    }
    return frequency
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
    frequency := wordFrequencyCount(input)
    fmt.Println(frequency)
}
