package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ueihvn/go-devduo/model"
)

type ProfileHandler struct {
	pr model.ProfileRepository
	tr model.TechnologyRepository
	fr model.FieldRepository
}

func NewProfileHandler(profileRepository model.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{
		pr: profileRepository,
	}
}

func (p *ProfileHandler) Create(w http.ResponseWriter, r *http.Request) {
	// get profilejson
	var profileJSON model.ProfileJSON
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &profileJSON)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}

	// +validate json request

	// convert from profileJSON to profile
	profile, err := profileJSON.ToProfile()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("%+v\n", err)
		return
	}

	// create profile
	err = p.pr.Create(profile)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		// duplicate err
		// database err
		return
	}

	// respone profileid
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (p *ProfileHandler) Get(w http.ResponseWriter, r *http.Request) {
	//get profileId from Request
	// getProfile -> get profile's techs -> get profile's fields
	// return to user
}
