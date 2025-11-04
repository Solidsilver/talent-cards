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
	talents := make(chan types.TalentSchema, 100)

	for range 10 {
		wg.Add(1)
		go func(tchan <-chan types.TalentSchema, wg *sync.WaitGroup) {
			defer wg.Done()
			for t := range tchan {
				fmt.Printf("  - Generating card for: %s\n", t.Name)
				err := card.CreateTalentCardHTML(sourcebook.Meta.Source.Full, t, outputDir)
				if err != nil {
					log.Printf("Could not create card for %s: %v", t.Name, err)
				}
			}
		}(talents, &wg)
	}

	for _, talent := range sourcebook.Talent {
		talents <- talent
	}
	close(talents)
	wg.Wait()

	card.Cleanup()

	fmt.Println("\nCard generation complete!")
	fmt.Printf("Cards saved in: %s\n", outputDir)
}

func TalentWorker(wg *sync.WaitGroup, talent chan *types.TalentSchema) {
	wg.Add(1)

}
