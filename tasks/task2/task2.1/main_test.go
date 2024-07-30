package main

import (
    "reflect"
    "testing"
)

func TestWordFrequencyCount(t *testing.T) {
    input := "I am Learning Go for internship, interview preparation and for fun"
    expected := map[string]int{
		"am":1,
		"and":1,
		"for":2,
		"fun":1,
		"go":1,
		"i":1,
		"internship":1,
		"interview":1,
		"learning":1,
		"preparation":1,
    }

    result := wordFrequencyCount(input)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("wordFrequencyCount(%q) = %v; want %v", input, result, expected)
    }
}
