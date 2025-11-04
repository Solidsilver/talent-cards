package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"talent-cards/pkg/card"
	"talent-cards/pkg/types"

	"github.com/chromedp/chromedp"
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

	numCPU := runtime.NumCPU()
	for id := range numCPU {
		wg.Go(func() {
			defer wg.Done()
			// Create unique allocator per worker to avoid singleton lock issues
			allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(),
				chromedp.ExecPath(card.GetExecPath()),
				chromedp.Headless,
				chromedp.NoDefaultBrowserCheck,
				chromedp.NoFirstRun,
				chromedp.UserDataDir(fmt.Sprintf("/tmp/chromedp-cache-%d", id)),
			)
			defer allocCancel()
			for t := range talents {
				fmt.Printf("  - Generating card for: %s\n", t.Name)
				err := card.CreateTalentCardHTMLWithAllocator(allocCtx, sourcebook.Meta.Source.Full, t, outputDir)
				if err != nil {
					log.Printf("Could not create card for %s: %v", t.Name, err)
				}
			}
		})
	}

	for _, talent := range sourcebook.Talent {
		talents <- talent
	}
	close(talents)
	wg.Wait()

	fmt.Println("\nCard generation complete!")
	fmt.Printf("Cards saved in: %s\n", outputDir)
}
