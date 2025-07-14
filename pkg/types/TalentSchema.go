package types

type TalentSchema struct {
	Activation    TalentActivation    `json:"activation"`
	Description   []string            `json:"description"`
	ImageURL      *string             `json:"imageUrl,omitempty"`
	Name          string              `json:"name"`
	Page          *int64              `json:"page,omitempty"`
	Prerequisites *Prerequisites      `json:"prerequisites"`
	Ranked        bool                `json:"ranked"`
	Settings      []SettingsSubSchema `json:"settings,omitempty"`
	Tags          []ItemTagsSubSchema `json:"tags,omitempty"`
	Tier          int64               `json:"tier"`
}
