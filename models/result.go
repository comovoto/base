package models

// Result Containt all info about election by candidate
type Result struct {
	Name           string `json:"name"`
	District       string `json:"district"`
	VotesList      string `json:"votes_list"`
	PercentageList string `json:"percentage_list"`
	Votes          string `json:"votes"`
	Percentage     string `json:"percentage"`
	Elect          bool   `json:"elect"`
}
