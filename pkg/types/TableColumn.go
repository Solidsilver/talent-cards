package types

type TableColumn struct {
	Header string   `json:"header"`
	Size   *float64 `json:"size,omitempty"`
	Value  *string  `json:"value,omitempty"`
}
