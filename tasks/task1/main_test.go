// tasks/main_test.go
package main

import (
    "testing"
)

func TestCalculateAverage(t *testing.T) {
    grades := []float64{85, 90, 80}
    expected := 85.0
    result := calculateAverage(grades)
    if result != expected {
        t.Errorf("calculateAverage(%v) = %v; want %v", grades, result, expected)
    }

    grades = []float64{70, 75, 80}
    expected = 75.0
    result = calculateAverage(grades)
    if result != expected {
        t.Errorf("calculateAverage(%v) = %v; want %v", grades, result, expected)
    }



    grades = []float64{100}
    expected = 100.0
    result = calculateAverage(grades)
    if result != expected {
        t.Errorf("calculateAverage(%v) = %v; want %v", grades, result, expected)
    }
}
