package types

type Meta struct {
	DefaultItemSettings []SettingsSubSchema `json:"defaultItemSettings"`
	Filters             *Filters            `json:"filters,omitempty"`
	Source              Source              `json:"source"`
}
