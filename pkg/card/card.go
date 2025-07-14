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

func parseActivation(activation types.TalentActivation) template.HTML {
	switch activation {
	case types.ActiveAction:
		return template.HTML("Action")
	case types.ActiveIncidental:
		return template.HTML("Incidental")
	case types.ActiveIncidentalOutOfTurn:
		return template.HTML("Incidental<br/>(Out of Turn)")
	case types.ActiveManeuver:
		return template.HTML("Maneuver")
	default:
		return template.HTML("<span style=\"font-weight: normal;\">Passive</span>")
	}
}

// CreateTalentCardHTML generates a PNG image for a talent by rendering an HTML template.
func CreateTalentCardHTML(source string, talent types.TalentSchema, outputDir string) error {
	// 1. Prepare data for the template
	data := CardTemplateData{
		Name:        talent.Name,
		Tier:        talent.Tier,
		Ranked:      talent.Ranked,
		Activation:  parseActivation(talent.Activation),
		Description: parser.ParseDescriptionToHTML(talent.Description),
		Page:        talent.Page,
		Source:      source,
	}

	// 2. Read and execute the HTML template
	templateBytes, err := os.ReadFile("assets/card-template.html")
	if err != nil {
		return fmt.Errorf("failed to read HTML template: %w", err)
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

	parentContext, parentContextCancel := chromedp.NewExecAllocator(context.Background(),
		chromedp.ExecPath(getExecPath()),
		chromedp.Headless,
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
		chromedp.CombinedOutput(os.Stdout),
	)
	defer parentContextCancel()
	ctx, cancel := chromedp.NewContext(parentContext)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var buf []byte
	fileURL := "file://" + filePath
	err = chromedp.Run(ctx,
		chromedp.Navigate(fileURL),
		chromedp.WaitVisible(`#card`, chromedp.ByID),
		// chromedp.Sleep(120*time.Millisecond),
		chromedp.Screenshot(`#card`, &buf, chromedp.ByID),
	)

	if err != nil {
		return fmt.Errorf("failed to take screenshot: %w", err)
	}

	safeFilename := strings.ReplaceAll(talent.Name, "/", "-")
	outputPath := filepath.Join(outputDir, fmt.Sprintf("%s.png", safeFilename))
	if err := os.WriteFile(outputPath, buf, 0644); err != nil {
		return fmt.Errorf("failed to save PNG: %w", err)
	}

	return nil
}

func getExecPath() string {
	switch runtime.GOOS {
	case "windows":
		return `%ProgramFiles%\Google\Chrome\Application\chrome.exe`
	case "darwin":
		return `/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`
	default:
		return "google-chrome"
	}
}
