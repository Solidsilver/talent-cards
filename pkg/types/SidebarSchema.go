package types

type SidebarSchema struct {
	Description []ChapterElement `json:"description"`
	Name        string           `json:"name"`
	Page        *int64           `json:"page,omitempty"`
	Summary     []Damage         `json:"summary,omitempty"`
	Type        string           `json:"type"`
}
