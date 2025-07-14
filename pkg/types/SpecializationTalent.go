package types

type SpecializationTalent struct {
	Down   *bool   `json:"down,omitempty"`
	Name   string  `json:"name"`
	Right  *bool   `json:"right,omitempty"`
	Source *string `json:"source,omitempty"`
}
