package types

type CharacteristicSchema struct {
	Description []string `json:"description"`
	Name        string   `json:"name"`
	Page        *int64   `json:"page,omitempty"`
}
