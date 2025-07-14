package types

type SpellSchema struct {
	AdditionalEffects []AdditionalEffect `json:"additionalEffects"`
	Concentration     bool               `json:"concentration"`
	Difficulty        Difficulty         `json:"difficulty"`
	ImageURL          *string            `json:"imageUrl,omitempty"`
	Name              string             `json:"name"`
	Narrative         []string           `json:"narrative"`
	Page              *int64             `json:"page,omitempty"`
	Range             *SpellRange        `json:"range,omitempty"`
	Skills            []SpellSkill       `json:"skills"`
	Structured        []string           `json:"structured"`
}
