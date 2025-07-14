package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"talent-cards/pkg/card"
	"talent-cards/pkg/types"
)

func main() {
	var jsonPath string
	var outputDir string
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <path_to_sourcebook.json> <output_directory>")
		os.Exit(1)
	} else {
		jsonPath = os.Args[1]
		outputDir = os.Args[2]
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	file, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var sourcebook types.Sourcebook
	if err := json.Unmarshal(file, &sourcebook); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Printf("Found %d talents. Generating cards using HTML template...\n", len(sourcebook.Talent))

	// Use a WaitGroup to process cards in parallel for speed.
	var wg sync.WaitGroup
	// Use a semaphore to limit concurrent browser instances to avoid overwhelming your system.
	sem := make(chan struct{}, 8) // Limit to 8 concurrent jobs

	for _, talent := range sourcebook.Talent {
		wg.Add(1)
		sem <- struct{}{} // Acquire a spot

		go func(t types.TalentSchema) {
			defer wg.Done()
			defer func() { <-sem }() // Release the spot

			fmt.Printf("  - Generating card for: %s\n", t.Name)
			// *** THIS IS THE ONLY LINE THAT REALLY CHANGED ***
			err := card.CreateTalentCardHTML(sourcebook.Meta.Source.Full, t, outputDir)
			if err != nil {
				// Log errors instead of crashing the whole process.
				log.Printf("Could not create card for %s: %v", t.Name, err)
			}
		}(talent)
	}

	wg.Wait() // Wait for all goroutines to finish.
	close(sem)

	fmt.Println("\nCard generation complete!")
	fmt.Printf("Cards saved in: %s\n", outputDir)
}
