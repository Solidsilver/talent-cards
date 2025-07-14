package types

type SpellEffectsSchema struct {
	Add    []Add    `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
	Spell  string   `json:"spell"`
}
