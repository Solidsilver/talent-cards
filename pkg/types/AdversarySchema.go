package types

type AdversarySchema struct {
	Abilities       []AdversaryAbility       `json:"abilities,omitempty"`
	Characteristics AdversaryCharacteristics `json:"characteristics"`
	Derived         DerivedClass             `json:"derived"`
	Description     []ChapterElement         `json:"description,omitempty"`
	Gear            []string                 `json:"gear,omitempty"`
	HasVariants     *bool                    `json:"hasVariants,omitempty"`
	Motivations     []MotivationElement      `json:"motivations,omitempty"`
	Name            string                   `json:"name"`
	Npc             *bool                    `json:"npc,omitempty"`
	Page            *int64                   `json:"page,omitempty"`
	PowerLevels     *PowerLevels             `json:"powerLevels,omitempty"`
	Settings        []SettingsSubSchema      `json:"settings,omitempty"`
	Skills          []AdversarySkill         `json:"skills"`
	Spells          *Spells                  `json:"spells,omitempty"`
	Tags            []string                 `json:"tags,omitempty"`
	Talents         []TalentUnion            `json:"talents,omitempty"`
	Type            AdversaryType            `json:"type"`
	Variant         *bool                    `json:"variant,omitempty"`
	Weapons         []AdversaryWeapon        `json:"weapons,omitempty"`
}
