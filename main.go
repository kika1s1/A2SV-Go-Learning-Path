package main

import (
    "fmt"
    "log"
)

// Function to calculate the average grade
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

    // Prompt user for their name
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    // Prompt user for the number of subjects
    fmt.Print("Enter the number of subjects: ")
    _, err := fmt.Scanln(&numSubjects)
    if err != nil || numSubjects <= 0 {
        log.Fatalf("Invalid number of subjects: %v", err)
    }

    // Create slices to store subject names and grades
    subjects := make([]string, numSubjects)
    grades := make([]float64, numSubjects)

    // Collect subject names and grades
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

    // Calculate the average grade
    average := calculateAverage(grades)

    // Display the results
    fmt.Printf("\nStudent Name: %s\n", name)
    fmt.Println("Subject Grades:")
    for i := 0; i < numSubjects; i++ {
        fmt.Printf("  %s: %.2f\n", subjects[i], grades[i])
    }
    fmt.Printf("Average Grade: %.2f\n", average)
}
