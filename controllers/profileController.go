package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	u "happyFeet.profileAI/utils"

	"happyFeet.profileAI/models"
)

// GetProfile response to a user request
var GetProfile = func(w http.ResponseWriter, r *http.Request) {

	//userID := r.Context.Value("userId").(string)
	params := mux.Vars(r)
	userID, err := params["userId"]
	if err != nil {

	}
//	profile := &models.Profile{}

	err := json.NewDecoder(r.Body).Decode(profile)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	profile.UserID = userID
	// resp :=
	u.Respond(w, resp)

}
