package models

import u "github.com/sisimogangg/happyFeet.profileAI/utils"

//Address stores users address info
type Address struct {
	StreetNumber int32  `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	PostalCode   int32  `json:"postalCode"`
}

// Validate validates the address
func (a *Address) Validate() (map[string]interface{}, bool) {
	if a.StreetNumber == 0 || len(a.StreetName) == 0 || len(a.Suburb) == 0 || len(a.City) == 0 || a.PostalCode == 0 {
		return u.Message(false, "Missing fields"), false
	}

	return u.Message(false, "Requirements Passed"), true
}
