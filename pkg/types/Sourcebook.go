// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    sourcebook, err := UnmarshalSourcebook(bytes)
//    bytes, err = sourcebook.Marshal()

package types

import "encoding/json"

func UnmarshalSourcebook(data []byte) (Sourcebook, error) {
	var r Sourcebook
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Sourcebook) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Sourcebook struct {
	Schema           string                 `json:"$schema"`
	Meta             Meta                   `json:"_meta"`
	Adventure        []AdventureSchema      `json:"adventure,omitempty"`
	Adversary        []AdversarySchema      `json:"adversary,omitempty"`
	AdversaryAbility []AbilitySchema        `json:"adversaryAbility,omitempty"`
	Archetype        []ArchetypeSchema      `json:"archetype,omitempty"`
	ArchetypeAbility []AbilitySchema        `json:"archetypeAbility,omitempty"`
	Book             []BookSchema           `json:"book,omitempty"`
	Career           []CareerSchema         `json:"career,omitempty"`
	Characteristic   []CharacteristicSchema `json:"characteristic,omitempty"`
	Gear             []GearSchema           `json:"gear,omitempty"`
	OptionFeature    []OptionFeatureSchema  `json:"optionFeature,omitempty"`
	Quality          []QualitySchema        `json:"quality,omitempty"`
	Rule             []RuleSchema           `json:"rule,omitempty"`
	Setting          []SettingSchema        `json:"setting,omitempty"`
	Sidebar          []SidebarSchema        `json:"sidebar,omitempty"`
	Skill            []SkillSchema          `json:"skill,omitempty"`
	Specialization   []SpecializationSchema `json:"specialization,omitempty"`
	Spell            []SpellSchema          `json:"spell,omitempty"`
	SpellEffects     []SpellEffectsSchema   `json:"spellEffects,omitempty"`
	Table            []TableSchema          `json:"table,omitempty"`
	Talent           []TalentSchema         `json:"talent,omitempty"`
	Vehicle          []VehicleSchema        `json:"vehicle,omitempty"`
}
