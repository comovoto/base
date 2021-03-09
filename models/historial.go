package models

// Historical ...
type Historical struct {
	ID                       string `json:"id"`
	URI                      string `json:"uri"`
	Year                     string `json:"year"`
	Election                 string `json:"election"`
	Territory                string `json:"territory"`
	IsElected                string `json:"is_elected"`
	List                     string `json:"list"`
	PoliticalParty           string `json:"political_party"`
	TotalVotesCandidate      string `json:"total_votes_candidate"`
	TotalVotesTerritory      string `json:"total_votes_territory"`
	PercentageVotesCandidate string `json:"percentage_votes_candidate"`
}
