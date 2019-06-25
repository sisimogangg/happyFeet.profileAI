package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	u "happyFeet.profileAI/utils"

	"github.com/gorilla/mux"
	"happyFeet.profileAI/models"
)

var GetProfile = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	userId := r.Context.Value("userid").(string)
	profile := &models.Profile{}

	err := json.NewDecoder(r.Body).Decode(profile)
	if err != nil {

	}

}
