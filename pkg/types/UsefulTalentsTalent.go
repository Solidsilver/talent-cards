package types

type UsefulTalentsTalent struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
	Tier   int64   `json:"tier"`
}
