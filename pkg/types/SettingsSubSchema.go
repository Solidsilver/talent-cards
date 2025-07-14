package types

type SettingsSubSchema struct {
	Name   string  `json:"name"`
	Source *string `json:"source,omitempty"`
}
