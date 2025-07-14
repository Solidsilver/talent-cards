package types

type TableSchema struct {
	Columns  []TableColumn `json:"columns"`
	Foot     *string       `json:"foot,omitempty"`
	Name     string        `json:"name"`
	Page     *int64        `json:"page,omitempty"`
	Rows     []TableRow    `json:"rows"`
	Subtitle *string       `json:"subtitle,omitempty"`
}
