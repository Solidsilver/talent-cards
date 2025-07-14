package types

type GearQuality struct {
	Name   string  `json:"name"`
	Ranks  *Damage `json:"ranks"`
	Source *string `json:"source,omitempty"`
}
