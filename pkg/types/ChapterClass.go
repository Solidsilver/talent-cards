package types

type ChapterClass struct {
	Ref             *Ref           `json:"_ref,omitempty"`
	Entries         []EntryElement `json:"entries,omitempty"`
	ForceTitleLevel *int64         `json:"forceTitleLevel,omitempty"`
	HasModal        *bool          `json:"hasModal,omitempty"`
	Page            *int64         `json:"page,omitempty"`
	Title           *string        `json:"title,omitempty"`
	Type            *ChapterType   `json:"type,omitempty"`
	Variant         *Variant       `json:"variant,omitempty"`
	By              *string        `json:"by,omitempty"`
	Description     *Description   `json:"description"`
	Items           []interface{}  `json:"items,omitempty"`
	Columns         *Columns       `json:"columns"`
	Foot            *string        `json:"foot,omitempty"`
	Rows            []ChapterRow   `json:"rows,omitempty"`
	Subtitle        *string        `json:"subtitle,omitempty"`
	ImageURL        *string        `json:"imageUrl,omitempty"`
}
