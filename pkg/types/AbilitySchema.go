package types

type AbilitySchema struct {
	// Ability details                                   
	Description                      []ChapterElement    `json:"description"`
	// Ability name.                                     
	Name                             string              `json:"name"`
	// Applicable tags for filtering.                    
	Tags                             []ItemTagsSubSchema `json:"tags"`
}
