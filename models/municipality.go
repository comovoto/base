package models

// Municipality ...
type Municipality struct {
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	DistrictID                string `json:"district_id"`
	DistrictName              string `json:"district_name"`
	ProvincialCircumscription string `json:"provincial_circumscription"`
	NumberElected             string `json:"number_elected"`
}
