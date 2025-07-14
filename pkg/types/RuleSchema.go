package types

type RuleSchema struct {
	Description []ChapterElement `json:"description"`
	Name        string           `json:"name"`
	Page        *int64           `json:"page,omitempty"`
	Type        string           `json:"type"`
}
