package types

type AdventureSchema struct {
	Chapters []ChapterElement    `json:"chapters"`
	Name     string              `json:"name"`
	Settings []SettingsSubSchema `json:"settings,omitempty"`
}
