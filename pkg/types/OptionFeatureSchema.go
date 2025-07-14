package types

type OptionFeatureSchema struct {
	Class       *string                `json:"class,omitempty"`
	Description []ChapterElement       `json:"description"`
	Info        []ChapterElement       `json:"info,omitempty"`
	Labels      []Label                `json:"labels,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
	Name        string                 `json:"name"`
	Page        *int64                 `json:"page,omitempty"`
	Settings    []SettingsSubSchema    `json:"settings,omitempty"`
	Summary     []string               `json:"summary,omitempty"`
	Type        string                 `json:"type"`
}
