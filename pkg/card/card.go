package card

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"talent-cards/pkg/parser"
	"talent-cards/pkg/types"
	"time"

	"github.com/chromedp/chromedp"
)

// CardTemplateData holds the data to be injected into the HTML template.
type CardTemplateData struct {
	Name        string
	Tier        int64
	Ranked      bool
	Activation  template.HTML
	Page        *int64
	Description template.HTML // Use template.HTML to prevent escaping of our generated HTML
	Source      string
}

var templateBytes []byte

func init() {
	var err error
	templateBytes, err = os.ReadFile("assets/card-template.html")
	if err != nil {
		panic(fmt.Errorf("failed to read HTML template: %w", err))
	}
}

// CreateTalentCardHTMLWithAllocator generates a PNG image for a talent by rendering an HTML template using the provided allocator.
func CreateTalentCardHTMLWithAllocator(allocCtx context.Context, source string, talent types.TalentSchema, outputDir string) error {
	// 1. Prepare data for the template
	data := CardTemplateData{
		Name:        talent.Name,
		Tier:        talent.Tier,
		Ranked:      talent.Ranked,
		Activation:  parser.ParseActivation(talent.Activation),
		Description: parser.ParseDescriptionToHTML(talent.Description),
		Page:        talent.Page,
		Source:      source,
	}

	// 2. Read and execute the HTML template
	// templateBytes, err := os.ReadFile("assets/card-template.html")
	if templateBytes == nil {
		return fmt.Errorf("failed to read HTML template")
	}

	tmpl, err := template.New("card").Parse(string(templateBytes))
	if err != nil {
		return fmt.Errorf("failed to parse HTML template: %w", err)
	}

	var processedHTML bytes.Buffer
	if err := tmpl.Execute(&processedHTML, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	tmpFile, err := os.CreateTemp(outputDir, "talent-card-*.html")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temp file

	if _, err := tmpFile.Write(processedHTML.Bytes()); err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}
	filePath, err := filepath.Abs(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var screenshotBuf []byte
	fileURL := "file://" + filePath
	err = chromedp.Run(ctx,
		chromedp.Navigate(fileURL),
		chromedp.WaitVisible(`#card`, chromedp.ByID),
		chromedp.Sleep(1*time.Second), // Increased sleep to ensure all external resources (CSS, fonts, images) are loaded and applied
		chromedp.Screenshot(`#card`, &screenshotBuf, chromedp.ByID),
	)

	if err != nil {
		return fmt.Errorf("failed to take screenshot: %w", err)
	}

	safeFilename := strings.ReplaceAll(talent.Name, "/", "-")
	outputPath := filepath.Join(outputDir, fmt.Sprintf("%s.png", safeFilename))
	if err := os.WriteFile(outputPath, screenshotBuf, 0644); err != nil {
		return fmt.Errorf("failed to save PNG: %w", err)
	}

	return nil
}

func GetExecPath() string {
	switch runtime.GOOS {
	case "windows":
		return `%ProgramFiles%\Google\Chrome\Application\chrome.exe`
	case "darwin":
		return `/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`
	default:
		return "google-chrome"
	}
}
