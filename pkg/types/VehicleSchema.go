package types

type VehicleSchema struct {
	Abilities           []VehicleAbility    `json:"abilities,omitempty"`
	Armor               int64               `json:"armor"`
	Complement          string              `json:"complement"`
	Consumables         string              `json:"consumables"`
	ControlSkill        ControlSkill        `json:"controlSkill"`
	Defense             int64               `json:"defense"`
	Description         []ChapterElement    `json:"description"`
	EncumbranceCapacity *Damage             `json:"encumbranceCapacity"`
	FtlRange            *FtlRange           `json:"ftlRange,omitempty"`
	FtlSpeed            *FtlSpeed           `json:"ftlSpeed,omitempty"`
	Handling            int64               `json:"handling"`
	HasVariants         *bool               `json:"hasVariants,omitempty"`
	Htt                 int64               `json:"htt"`
	ImageURL            *string             `json:"imageUrl,omitempty"`
	MaxSpeed            int64               `json:"maxSpeed"`
	Name                string              `json:"name"`
	Page                *int64              `json:"page,omitempty"`
	Passengers          *Damage             `json:"passengers"`
	Price               *int64              `json:"price,omitempty"`
	Rarity              int64               `json:"rarity"`
	Restricted          *bool               `json:"restricted,omitempty"`
	Settings            []SettingsSubSchema `json:"settings,omitempty"`
	Silhouette          int64               `json:"silhouette"`
	Sst                 int64               `json:"sst"`
	Variant             *bool               `json:"variant,omitempty"`
	Weapons             []VehicleWeapon     `json:"weapons,omitempty"`
}
