package controllers

import (
	"net/http"
	"time"

	"happyFeet.profileAI/models"

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
	kid := models.Kid{Name: "Thubelihle", Grade: 0, DateStarted: time.Now()}
	address := models.Address{StreetNumber: 16, StreetName: "Lonsdale Drive", Suburb: "Durban North", City: "Durban", PostalCode: 4051}
	personalDets := models.PersonalDetails{Name: "Godfrey Sisimogang", Address: address}
	err := personalDets.SetDateOfBirth("1990-03-10")
	if err != nil {
		panic(err)
	}

	kids := make([]models.Kid, 0, 1)
	kids = append(kids, kid)

	profile := models.Profile{UserID: userID, Kids: kids, PersonalDetails: personalDets}

	resp := u.Message(true, "success")
	resp["profile"] = profile

	u.Respond(w, resp)

}
