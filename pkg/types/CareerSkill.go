package types

type CareerSkill struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
