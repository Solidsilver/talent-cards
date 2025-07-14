package types

type SkillSchema struct {
	Category       Category            `json:"category"`
	Characteristic string              `json:"characteristic"`
	Description    []ChapterElement    `json:"description"`
	ImageURL       *string             `json:"imageUrl,omitempty"`
	Name           string              `json:"name"`
	Page           *int64              `json:"page,omitempty"`
	Settings       []SettingsSubSchema `json:"settings,omitempty"`
	ShouldNotUse   []string            `json:"shouldNotUse"`
	ShouldUse      []string            `json:"shouldUse"`
	Spells         []SkillSpell        `json:"spells,omitempty"`
}
