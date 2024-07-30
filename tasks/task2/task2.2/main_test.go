package main

import (
    "testing"
)


func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"this, is not palandrome", false},
        {"My name is Tamirat", false},
        {"adadada", true},
        
    }

    for _, test := range tests {
        result := isPalindrome(test.input)
        if result != test.expected {
            t.Errorf("isPalindrome(%q) = %v; want %v", test.input, result, test.expected)
        }
    }
}
