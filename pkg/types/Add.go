package types

type Add struct {
	Description   []string `json:"description"`
	DifficultyMod int64    `json:"difficultyMod"`
	Name          string   `json:"name"`
}
