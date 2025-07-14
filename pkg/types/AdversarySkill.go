package types

type AdversarySkill struct {
	Characteristic CharacteristicEnum `json:"characteristic"`
	Name           string             `json:"name"`
	Ranks          *int64             `json:"ranks,omitempty"`
	Source         *string            `json:"source,omitempty"`
}
