package types

type WeaponQualityClass struct {
	Name   string  `json:"name"`
	Ranks  *int64  `json:"ranks,omitempty"`
	Source *string `json:"source,omitempty"`
}
