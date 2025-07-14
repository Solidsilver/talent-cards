package types

type WeaponSkillClass struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
