package types

type BookSchema struct {
	BackgroundURL      *string                `json:"backgroundUrl,omitempty"`
	Chapters           []ChapterElement       `json:"chapters"`
	CoverBackgroundURL *string                `json:"coverBackgroundUrl,omitempty"`
	CoverImageURL      *string                `json:"coverImageUrl,omitempty"`
	Credits            map[string]interface{} `json:"credits,omitempty"`
	Name               string                 `json:"name"`
	ReleaseYear        *int64                 `json:"releaseYear,omitempty"`
	Subtitle           *string                `json:"subtitle,omitempty"`
}
