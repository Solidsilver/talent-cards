package types

type ArchetypeSchema struct {
	Abilities       []ArchetypeAbility    `json:"abilities"`
	Characteristics *CharacteristicsUnion `json:"characteristics"`
	Description     []string              `json:"description"`
	ImageURL        *string               `json:"imageUrl,omitempty"`
	Name            string                `json:"name"`
	Names           *Names                `json:"names,omitempty"`
	Nickname        *string               `json:"nickname,omitempty"`
	Page            *int64                `json:"page,omitempty"`
	Settings        []SettingsSubSchema   `json:"settings,omitempty"`
	Skills          Skills                `json:"skills"`
	St              *St                   `json:"st"`
	Why             []string              `json:"why,omitempty"`
	Wt              *St                   `json:"wt"`
	XP              int64                 `json:"xp"`
}
