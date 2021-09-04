package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ueihvn/go-devduo/model"
)

type ProfileHandler struct {
	pr model.ProfileRepository
}

func NewProfileHandler(profileRepository model.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{
		pr: profileRepository,
	}
}

// swagger:route POST /api/v1/profile profile profileCreateUpdateReq
// create profile
// responses:
// 200: profileCreateUpdateResp
// 400: errorResp
// 500: errorResp
func (p *ProfileHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get profilejson
	var profileJSON model.ProfileJSON
	err := FromJSON(&profileJSON, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
		return
	}

	// +validate json request

	// convert from profileJSON to profile
	profile, err := profileJSON.ToProfile()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
		return
	}

	// create profile
	err = p.pr.Create(profile)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			err = errors.New("profile already exits with userid")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err = errors.New("database error")
			w.WriteHeader(http.StatusInternalServerError)
		}
		ToJSON(Response{Status: false, Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully create profile", Data: profileJSON}, w)
}

// swagger:operation GET /api/v1/profile/{id} profile get
// ---
// summary: Get user profile by user id
// parameters:
// - name: id
//   in: path
//   description: id of profile
//   required: true
//   type: integer
//   format: unit64
// responses:
//   "200":
//     "$ref": "#/responses/profileCreateUpdateResp"
//   "403":
//     "$ref": "#/responses/errorResp"
//   "500":
//     "$ref": "#/responses/errorResp"
func (p *ProfileHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID, err := parseID(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: ParseIDError}, w)
		return
	}

	profile, err := p.pr.Get(userID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			ToJSON(Response{Status: false, Message: notFoundError}, w)
		}
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: serverInternalError}, w)
		return
	}

	profileJSON, err := profile.ToProfileJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: serverInternalError}, w)
		return
	}

	ToJSON(Response{Status: true, Message: "successfully ger profile", Data: profileJSON}, w)
}

// swagger:route PUT /api/v1/profile profile profileCreateUpdateReq
// update profile
// responses:
// 200: profileCreateUpdateResp
// 400: errorResp
// 500: errorResp
func (p *ProfileHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var profileJSON model.ProfileJSON
	err := FromJSON(&profileJSON, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
		return
	}

	// +validate json request

	// convert from profileJSON to profile
	profile, err := profileJSON.ToProfile()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
		return
	}

	err = p.pr.Update(profile)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			ToJSON(Response{Status: false, Message: notFoundError}, w)
		}
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: serverInternalError}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully update profile", Data: profileJSON}, w)
}
