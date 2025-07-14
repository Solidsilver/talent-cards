package types

type UsefulTalents struct {
	Description []string              `json:"description"`
	Talents     []UsefulTalentsTalent `json:"talents"`
}
