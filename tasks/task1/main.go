package main

import (
    "fmt"
    "log"
)

func calculateAverage(grades []float64) float64 {
    var total float64
    for _, grade := range grades {
        total += grade
    }
    return total / float64(len(grades))
}

func main() {
    var name string
    var numSubjects int

    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    fmt.Print("Enter the number of subjects: ")
    _, err := fmt.Scanln(&numSubjects)
    if err != nil || numSubjects <= 0 {
        log.Fatalf("Invalid number of subjects: %v", err)
    }

    subjects := make([]string, numSubjects)
    grades := make([]float64, numSubjects)

    for i := 0; i < numSubjects; i++ {
        fmt.Printf("Enter the name of subject %d: ", i+1)
        fmt.Scanln(&subjects[i])

        var grade float64
        fmt.Printf("Enter the grade for %s: ", subjects[i])
        _, err := fmt.Scanln(&grade)
        if err != nil || grade < 0 || grade > 100 {
            log.Fatalf("Invalid grade value: %v", err)
        }
        grades[i] = grade
    }

    average := calculateAverage(grades)

    fmt.Printf("\nStudent Name: %s\n", name)
    fmt.Println("Subject Grades:")
    for i := 0; i < numSubjects; i++ {
        fmt.Printf("  %s: %.2f\n", subjects[i], grades[i])
    }
    fmt.Printf("Average Grade: %.2f\n", average)
}
