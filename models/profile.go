package models

import u "github.com/sisimogangg/happyFeet.profileAI/utils"

// Profile stores profile info
type Profile struct {
	ID              int64           `json:"id"`
	UserID          string          `json:"userId"`
	Kids            []Kid           `json:"kids"`
	SGBMember       bool            `json:"sgbMember"`
	PersonalDetails PersonalDetails `json:"personalDetails"`
}

// GetProfile This Gets users profile from DB
func GetProfile(userID string) *Profile {
	profile := Profile{}

	profile.UserID = userID
	return &profile
}

// Validate validates profile
func (p *Profile) Validate() (map[string]interface{}, bool) {
	if p.UserID == "" {
		return u.Message(false, "User ID is required"), false
	}

	return u.Message(true, "Success"), true
}

// CreateProfile returns users profile
func (p *Profile) CreateProfile() map[string]interface{} {
	if resp, ok := p.Validate(); !ok {
		return resp
	}

	resp := u.Message(true, "Success")
	resp["profile"] = p
	return resp
}
