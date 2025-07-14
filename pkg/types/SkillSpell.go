package types

type SkillSpell struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
