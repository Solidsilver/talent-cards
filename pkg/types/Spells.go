package types

type Spells struct {
	Skills []SpellsSkill `json:"skills"`
	Spells []SpellsSpell `json:"spells"`
}
