package main

import "fmt"

func showHistory() {
	history := loadHistory()

	if len(history) == 0 {
		fmt.Println("No history available.")
		return
	}

	for i, h := range history {
		fmt.Printf("%d. %s\n", i+1, h.Question)
		fmt.Printf("   → %s\n\n", h.Command)
	}
}
