package main

import (
	"fmt"
	"github.com/kelbwah/pq-374/pq"
)

type ConditionLevel int

const (
    Mild ConditionLevel = iota + 1 // Lowest Priority
    Moderate 
    Severe
    Critical // Highest Priority
)


// Real life scenario: ER Triage
func main() {
    // Create a new Priority Queue
    taskQueue := pq.CreatePriorityQueue()
    conditionToString := map[ConditionLevel]string{
        Mild: "Mild",
        Moderate: "Moderate",
        Severe: "Severe",
        Critical: "Critical",
    }

    // Patients list, the higher the number, the more severe their condition.
    patients := []struct {
        name     string
        priority ConditionLevel 
    }{
        {"John", Critical},
        {"Maria", Severe},
        {"Juan", Moderate},
        {"Darius", Moderate},
        {"Nathan", Mild},
        {"Jacob", Severe},
        {"Kelly", Mild},
        {"Robert", Critical},
        {"Esmeralda", Severe},
        {"Caleb", Mild},
    }

    // Enqueue patients into the priority queue
    fmt.Println("Enqueuing patients into the priority queue:")
    for _, patient := range patients {
        fmt.Printf("Adding patient %s with condition level: %s\n", patient.name, conditionToString[patient.priority])
        taskQueue.Enqueue(int(patient.priority))
    }

    // Peek at the most important patient condition out of all patients waiting in ER 
    maxPriority, err := taskQueue.Peek()
    if err == nil {
        fmt.Printf("\nMost important patient condition is: %s\n", conditionToString[ConditionLevel(maxPriority)])
    }

    // Triaging patients based on priority
    fmt.Println("\nTriaging patients in order of their condition level:")
    for !taskQueue.IsEmpty() {
        priority, err := taskQueue.Dequeue()
        if err == nil {
          fmt.Printf("Helping patient with condition: %s\n", conditionToString[ConditionLevel(priority)])
        }
    }

    fmt.Println("\nAll patients triaged!")
}
