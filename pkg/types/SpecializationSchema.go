package types

type SpecializationSchema struct {
	Career      *Career                  `json:"career,omitempty"`
	Description []string                 `json:"description,omitempty"`
	ImageURL    *string                  `json:"imageUrl,omitempty"`
	Name        string                   `json:"name"`
	Page        *int64                   `json:"page,omitempty"`
	Settings    []SettingsSubSchema      `json:"settings,omitempty"`
	Skills      []SpecializationSkill    `json:"skills,omitempty"`
	Talents     [][]SpecializationTalent `json:"talents"`
}
