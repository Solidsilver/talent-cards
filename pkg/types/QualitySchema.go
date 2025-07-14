package types

type QualitySchema struct {
	Activation  QualityActivation `json:"activation"`
	Description []ChapterElement  `json:"description"`
	ItemTypes   ItemTypes         `json:"itemTypes"`
	Name        string            `json:"name"`
	Page        *int64            `json:"page,omitempty"`
}
