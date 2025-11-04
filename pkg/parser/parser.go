package parser

import (
	"fmt"
	"html/template"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"talent-cards/pkg/types"
)

// Regex to find an "innermost" tag, i.e., one that contains no other tags.
// The content part `[^}{]*` is key, as it matches any sequence of characters
// that does not contain a '{' or '}', ensuring we only get the deepest tags.
var innermostTagRegex = regexp.MustCompile(`{@([a-zA-Z]+)(?:\s+([^}{]*))?}`)

type diceInfo struct {
	char  string
	class string
}

var diceMap = map[string]diceInfo{
	"boost":       {"j", "boost"},
	"blue":        {"j", "boost"},
	"b":           {"j", "boost"},
	"setback":     {"j", "setback"},
	"black":       {"j", "setback"},
	"k":           {"j", "setback"},
	"ability":     {"k", "ability"},
	"green":       {"k", "ability"},
	"g":           {"k", "ability"},
	"difficulty":  {"k", "difficulty"},
	"purple":      {"k", "difficulty"},
	"p":           {"k", "difficulty"},
	"proficiency": {"l", "proficiency"},
	"yellow":      {"l", "proficiency"},
	"y":           {"l", "proficiency"},
	"challenge":   {"l", "challenge"},
	"red":         {"l", "challenge"},
	"r":           {"l", "challenge"},
}

var difficultyMap = map[string]int{
	"simple":     0,
	"easy":       1,
	"average":    2,
	"hard":       3,
	"daunting":   4,
	"formidable": 5,
}

var powerMap = map[string]diceInfo{
	"combat":  {"c", "combat"},
	"general": {"g", "general"},
	"social":  {"p", "social"},
}

type Ref struct {
	refType string
	value   string
}

// processSingleTag takes a tag type and its content and returns the HTML replacement.
func processSingleTag(tagType, content string) (string, *Ref) {
	args := strings.Split(content, "|")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}

	switch tagType {
	case "i":
		return fmt.Sprintf("<i>%s</i>", args[0]), nil
	case "b":
		return fmt.Sprintf("<b>%s</b>", args[0]), nil
	case "s":
		return fmt.Sprintf("<s>%s</s>", args[0]), nil
	case "title":
		return fmt.Sprintf(`<span style="font-weight: bold; text-transform: uppercase;">%s</span>`, args[0]), nil
	case "table":
		tableInfo := strings.Split(args[0], ":")
		ref := Ref{
			refType: "table",
			value:   tableInfo[0],
		}
		return fmt.Sprintf(`<span class="table-ref">%s*</span>`, tableInfo[1]), &ref
	case "dice":
		return handleDiceTag(args), nil
	case "difficulty":
		return handleDifficultyTag(args), nil
	case "symbols":
		return fmt.Sprintf(`<span class="genesys">%s</span>`, args[0]), nil
	case "combat", "general", "social":
		return handlePowerTag(tagType, args), nil
	case "talent":
		return fmt.Sprintf(`<span class="talent">%s</span>`, args[0]), nil
	case "skill":
		return fmt.Sprintf(`<span class="skill">%s</span>`, args[0]), nil
	case "rule":
		return fmt.Sprintf(`<span class="rule">%s</span>`, args[0]), nil
	case "quality":
		return fmt.Sprintf(`<span class="quality">%s</span>`, args[0]), nil
	case "spell":
		return fmt.Sprintf(`<span class="spell">%s (%s)</span>`, args[0], args[len(args)-1]), nil
	case "adversary":
		return fmt.Sprintf(`<span class="adversary">%s</span>`, args[0]), nil
	case "gear":
		// return fmt.Sprintf(`<span class="gear">%s (%s)</span>`, args[0], args[len(args)-1]), nil
		return fmt.Sprintf(`<span class="gear">%s</span>`, args[0]), nil
	default:
		// If tag is not recognized, return the original text to avoid losing it.
		// If you are trying to add support for new tags, this is the place to debug
		if content == "" {
			return fmt.Sprintf("%s", tagType), nil
		}
		return fmt.Sprintf("%s %s", tagType, content), nil
	}
}

// parseLine iteratively finds and replaces the innermost tags until none are left.
func parseLine(line string) (string, []Ref) {
	refs := []Ref{}
	for {
		match := innermostTagRegex.FindStringSubmatch(line)
		if match == nil {
			break
		}

		fullMatchString := match[0]
		tagType := strings.ToLower(match[1])
		content := ""
		if len(match) > 2 {
			content = match[2]
		}

		replacement, ref := processSingleTag(tagType, content)
		if ref != nil && !slices.Contains(refs, *ref) {
			refs = append(refs, *ref)
		}

		line = strings.Replace(line, fullMatchString, replacement, 1)
	}
	return line, refs
}

// ParseDescriptionToHTML converts an array of description strings into a single,
// safe HTML string with custom syntax replaced by HTML tags.
func ParseDescriptionToHTML(description []string) template.HTML {
	var htmlStrings []string
	refs := []Ref{}
	for _, line := range description {
		htmlLine, lineRefs := parseLine(line)
		htmlStrings = append(htmlStrings, htmlLine)
		for _, ref := range lineRefs {
			if !slices.Contains(refs, ref) {
				refs = append(refs, ref)
			}
		}
	}

	finalHTML := strings.Join(htmlStrings, "<br><br>")
	for _, ref := range refs {
		switch ref.refType {
		case "table":
			finalHTML += fmt.Sprintf(`<span class="table-ref"> *[See Table %s] </span>`, ref.value)
			break

		}

	}
	return template.HTML(finalHTML)
}

func handleDiceTag(args []string) string {
	if len(args) == 0 || args[0] == "" {
		return ""
	}
	dieName := strings.ToLower(args[0])
	info, ok := diceMap[dieName]
	if !ok {
		return ""
	}

	quantity := 1
	if len(args) > 1 {
		if q, err := strconv.Atoi(args[1]); err == nil {
			quantity = q
		}
	}

	upgrades := 0
	if len(args) > 2 {
		if u, err := strconv.Atoi(args[2]); err == nil {
			upgrades = u
		}
	}

	// Handle upgrades (ability -> proficiency, difficulty -> challenge)
	if (info.class == "ability" || info.class == "difficulty") && upgrades > 0 && quantity >= upgrades {
		upgradedInfo := diceMap["proficiency"]
		if info.class == "difficulty" {
			upgradedInfo = diceMap["challenge"]
		}

		upgradedDice := strings.Repeat(upgradedInfo.char, upgrades)
		baseDice := strings.Repeat(info.char, quantity-upgrades)

		// Return two separate spans to handle the different colors
		return fmt.Sprintf(`<span class="genesys dice %s">%s</span><span class="genesys dice %s">%s</span>`,
			upgradedInfo.class, upgradedDice, info.class, baseDice)
	}

	return fmt.Sprintf(`<span class="genesys dice %s">%s</span>`, info.class, strings.Repeat(info.char, quantity))
}

func handleDifficultyTag(args []string) string {
	if len(args) == 0 || args[0] == "" {
		return ""
	}
	difficultyName := strings.ToLower(args[0])
	numDice, ok := difficultyMap[difficultyName]
	if !ok {
		return ""
	}

	skill := ""
	if len(args) > 1 && args[1] != "" {
		skill = " " + args[1]
	}

	upgrades := 0
	if len(args) > 2 {
		if u, err := strconv.Atoi(args[2]); err == nil {
			upgrades = u
		}
	}

	if numDice == 0 {
		return fmt.Sprintf("(-) %s", skill)
	}

	challengeInfo := diceMap["challenge"]
	difficultyInfo := diceMap["difficulty"]

	challengeDice := ""
	if upgrades > 0 {
		challengeDice = strings.Repeat(challengeInfo.char, upgrades)
	}
	difficultyDice := strings.Repeat(difficultyInfo.char, numDice-upgrades)

	diceHTML := fmt.Sprintf(`<span class="genesys dice %s">%s</span><span class="genesys dice %s">%s</span>`,
		challengeInfo.class, challengeDice, difficultyInfo.class, difficultyDice)

	return fmt.Sprintf(`%s (%s)%s`, strings.ToTitle(difficultyName), diceHTML, skill)
}

func handlePowerTag(tagType string, args []string) string {
	if len(args) == 0 || args[0] == "" {
		return ""
	}
	info, ok := powerMap[tagType]
	if !ok {
		return ""
	}
	return fmt.Sprintf(`<span class="genesys %s">%s</span>%s`, info.class, info.char, args[0])
}

func ParseActivation(activation types.TalentActivation) template.HTML {
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
