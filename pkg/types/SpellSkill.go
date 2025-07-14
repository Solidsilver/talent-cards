package types

type SpellSkill struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
