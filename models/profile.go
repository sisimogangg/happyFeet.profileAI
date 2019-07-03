package models

// Profile stores profile info
type Profile struct {
	UserID          string          `json:"userId"`
	Kids            []Kid           `json:"kids"`
	PersonalDetails PersonalDetails `json:"personalDetails"`
}

// GetProfile This Gets users profile from DB
func GetProfile(userID string) *Profile {
	profile := Profile{}

	profile.UserID = userID
	return &profile
}
