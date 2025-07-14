package types

type SpellsSkill struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
