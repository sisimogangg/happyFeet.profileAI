package models


type Profile struct {
	UserID string `json:"userId"`
	Kids  []Kid `json:"kids"`
	PersonalDetails PersonalDetails `json:"personalDetails"`    
}

func GetProfile(userId string)(*Profile){
	profile := Profile{}

	profile.UserID = userId
    return &profile
}


