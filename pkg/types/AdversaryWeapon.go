package types

type AdversaryWeapon struct {
	Critical  int64          `json:"critical"`
	Damage    int64          `json:"damage"`
	Details   *string        `json:"details,omitempty"`
	Name      string         `json:"name"`
	Qualities []QualityUnion `json:"qualities,omitempty"`
	Range     WeaponRange    `json:"range"`
	Skill     *SkillUnion    `json:"skill"`
}
