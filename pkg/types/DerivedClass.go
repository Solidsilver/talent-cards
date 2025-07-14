package types

type DerivedClass struct {
	Defense []int64 `json:"defense"`
	Soak    int64   `json:"soak"`
	Strain  *int64  `json:"strain,omitempty"`
	Wounds  int64   `json:"wounds"`
}
