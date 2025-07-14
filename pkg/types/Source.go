package types

type Source struct {
	Abbreviation                                                                              string   `json:"abbreviation"`
	Authors                                                                                   []string `json:"authors"`
	Color                                                                                     *string  `json:"color,omitempty"`
	ConvertedBy                                                                               []string `json:"convertedBy,omitempty"`
	Full                                                                                      string   `json:"full"`
	JSON                                                                                      string   `json:"json"`
	Language                                                                                  *string  `json:"language,omitempty"`
	// For official Modules. If true, Core Collection will not be installed by default for new         
	// users.                                                                                          
	Module                                                                                    *bool    `json:"module,omitempty"`
	ReleaseDate                                                                               *string  `json:"releaseDate,omitempty"`
	TranslatedBy                                                                              []string `json:"translatedBy,omitempty"`
	URL                                                                                       *string  `json:"url,omitempty"`
	Version                                                                                   string   `json:"version"`
}
