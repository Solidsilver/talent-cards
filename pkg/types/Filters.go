package types

type Filters struct {
	MagicSkills        []string `json:"magicSkills,omitempty"`
	OptionFeatureClass []string `json:"optionFeatureClass,omitempty"`
	OptionFeatureType  []string `json:"optionFeatureType,omitempty"`
	Settings           []string `json:"settings,omitempty"`
	TalentPrereqs      []string `json:"talentPrereqs,omitempty"`
}
