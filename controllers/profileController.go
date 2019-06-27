package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	u "happyFeet.profileAI/utils"
)

// GetProfile response to a user request
var GetProfile = func(w http.ResponseWriter, r *http.Request) {

	//userID := r.Context.Value("userId").(string)
	params := mux.Vars(r)
	userID := params["userId"]
	if userID == "" {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}
	//	profile := &models.Profile{}

	/*	err := json.NewDecoder(r.Body).Decode(profile)
		if err != nil {
			u.Respond(w, u.Message(false, "Error while decoding request body"))
			return
		} */

	//profile.UserID = userID
	// resp :=
	data := u.Message(true, "You're in baby")
	u.Respond(w, data)

}
