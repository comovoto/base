package models

// Candidate structure that contains the information of a candidate to the constituent process
type Candidate struct {
	Name            string         `json:"name"`
	URI             string         `json:"uri"`
	District        string         `json:"district"`
	List            string         `json:"list"`
	Age             string         `json:"age"`
	Quota           string         `json:"quota"`
	PoliticalParty  string         `json:"political_party"`
	ExtraInfo       string         `json:"extra_info"`
	Motivation      string         `json:"motivation"`
	Twitter         string         `json:"twitter"`
	Facebook        string         `json:"facebook"`
	Instagram       string         `json:"instagram"`
	Website         string         `json:"website"`
	Profession      string         `json:"profession"`
	Region          string         `json:"region"`
	Gender          string         `json:"gender"`
	IsIndependent   bool           `json:"is_independent"`
	ReservedSeat    string         `json:"reserved_seat"`
	OriginalTownsID string         `json:"original_towns_id"`
	OriginalTowns   string         `json:"original_towns"`
	URL             string         `json:"url"`
	Photo           string         `json:"photo"`
	Contributions   []Contribution `json:"contributions"`
	Votes           string         `json:"votes"`
	Percentage      string         `json:"percentage"`
	Elect           bool           `json:"elect"`
	VotesList       string         `json:"votes_list"`
	PercentageList  string         `json:"percentage_list"`
}
