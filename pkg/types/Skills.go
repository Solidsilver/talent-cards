package types

type Skills struct {
	Choice      bool     `json:"choice"`
	Description []string `json:"description"`
	RanksEach   int64    `json:"ranksEach"`
	RanksTotal  int64    `json:"ranksTotal"`
	Skills      []string `json:"skills"`
}
