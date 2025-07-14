package types

type GearSkill struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
