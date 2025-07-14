package types

type ControlSkill struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
