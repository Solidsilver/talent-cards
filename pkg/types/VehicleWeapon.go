package types

type VehicleWeapon struct {
	Critical    *int64               `json:"critical,omitempty"`
	Damage      *int64               `json:"damage,omitempty"`
	Description *string              `json:"description,omitempty"`
	Details     *string              `json:"details,omitempty"`
	FireArc     *string              `json:"fireArc,omitempty"`
	Name        string               `json:"name"`
	Qualities   []WeaponQualityClass `json:"qualities,omitempty"`
	Range       *WeaponRange         `json:"range,omitempty"`
	Skill       *WeaponSkillClass    `json:"skill,omitempty"`
}
