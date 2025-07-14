package types

type Career struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
