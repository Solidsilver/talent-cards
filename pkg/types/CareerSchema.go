package types

type CareerSchema struct {
	Description   []string            `json:"description"`
	Name          string              `json:"name"`
	Page          *int64              `json:"page,omitempty"`
	Settings      []SettingsSubSchema `json:"settings,omitempty"`
	Skills        []CareerSkill       `json:"skills"`
	StartingGear  *StartingGear       `json:"startingGear,omitempty"`
	UsefulTalents *UsefulTalents      `json:"usefulTalents,omitempty"`
}
