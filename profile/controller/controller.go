package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sisimogangg/happyFeet.profileAI/profile"
	u "github.com/sisimogangg/happyFeet.profileAI/utils"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ProfileHandler represent the httphandler for profile
type ProfileHandler struct {
	ProfileService profile.Servicer
}

// NewProfileHandler handles http requests
func NewProfileHandler(router *mux.Router, service profile.Servicer) {
	handler := &ProfileHandler{
		ProfileService: service,
	}

	router.HandleFunc("/api/profile/{userId}", handler.FetchProfile).Methods("GET")

}

// FetchProfile fetches profile from the service
func (h *ProfileHandler) FetchProfile(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	num, err := strconv.Atoi(params["num"])
	if err != nil {
		u.Respond(w, u.Message(false, "Limit must be a number"))
	}

	if num == 0 {
		u.Respond(w, u.Message(false, "There was an error in your request"))
	}

	ctx := r.Context()
	if ctx != nil {
		ctx = context.Background()
	}

	profiles, err := h.ProfileService.Fetch(ctx, int64(num))
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
	}

	resp := u.Message(true, "Success")
	resp["profiles"] = profiles

	u.Respond(w, resp)
}
