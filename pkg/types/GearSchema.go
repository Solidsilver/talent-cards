package types

type GearSchema struct {
	Class       *string             `json:"class,omitempty"`
	Critical    *int64              `json:"critical,omitempty"`
	Damage      *Damage             `json:"damage"`
	Description []EntryElement      `json:"description"`
	Encumbrance *Damage             `json:"encumbrance"`
	HardPoints  *HardPoints         `json:"hardPoints"`
	ImageURL    *string             `json:"imageUrl,omitempty"`
	Modifiers   []Modifier          `json:"modifiers,omitempty"`
	Name        string              `json:"name"`
	Page        *int64              `json:"page,omitempty"`
	Price       *PriceUnion         `json:"price"`
	Range       *WeaponRange        `json:"range,omitempty"`
	Rarity      *int64              `json:"rarity"`
	Restricted  *bool               `json:"restricted,omitempty"`
	Settings    []SettingsSubSchema `json:"settings,omitempty"`
	Skill       *GearSkill          `json:"skill,omitempty"`
	Special     []Special           `json:"special,omitempty"`
	Type        GearType            `json:"type"`
	Defense     *int64              `json:"defense,omitempty"`
	Soak        *string             `json:"soak,omitempty"`
	Qualities   []GearQuality       `json:"qualities,omitempty"`
}
