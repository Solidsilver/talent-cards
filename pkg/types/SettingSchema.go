package types

type SettingSchema struct {
	CharacterOptions []ChapterElement `json:"characterOptions,omitempty"`
	Codex            []ChapterElement `json:"codex,omitempty"`
	Lore             []ChapterElement `json:"lore,omitempty"`
	Name             string           `json:"name"`
	Page             *int64           `json:"page,omitempty"`
	Summary          []string         `json:"summary,omitempty"`
}
