package types

type AdditionalEffect struct {
	Description       []string `json:"description"`
	DifficultyMod     int64    `json:"difficultyMod"`
	DifficultyUpgrade *int64   `json:"difficultyUpgrade,omitempty"`
	Name              string   `json:"name"`
}
