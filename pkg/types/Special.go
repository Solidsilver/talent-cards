package types

type Special struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
	Value  *int64  `json:"value,omitempty"`
}
