package types

type Modifier struct {
	Name  string        `json:"name"`
	Type  *ModifierType `json:"type,omitempty"`
	Value *Damage       `json:"value"`
}
