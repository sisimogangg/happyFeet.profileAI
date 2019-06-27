package models

type Address struct {
	StreetNumber int32  `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	PostalCode   int32  `json:"postalCode"`
}
