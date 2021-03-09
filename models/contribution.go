package models

// Contribution structure that contains the information of a contribution by candidate
type Contribution struct {
	ID          string `json:"id"`
	Contributor string `json:"contributor"`
	Candidate   string `json:"candidate"`
	URI         string `json:"uri"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	Amount      string `json:"amount"`
}
