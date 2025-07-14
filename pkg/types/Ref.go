package types

type Ref struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
