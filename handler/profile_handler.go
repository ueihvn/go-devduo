package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ueihvn/go-devduo/model"
)

type ProfileHandler struct {
	pr model.ProfileRepository
	tr model.TechnologyRepository
	fr model.FieldRepository
}

func NewProfileHandler(profileRepository model.ProfileRepository, techRepository model.TechnologyRepository, fieldRepository model.FieldRepository) *ProfileHandler {
	return &ProfileHandler{
		pr: profileRepository,
		tr: techRepository,
		fr: fieldRepository,
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
	vars := mux.Vars(r)
	userID, err := parseID(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	profile, err := p.pr.Get(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	techs, err := p.tr.GetTechnologiesByUserId(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fields, err := p.fr.GetFieldsByUserId(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	profile.Technologies = techs
	profile.Fields = fields

	profileJSON, err := profile.ToProfileJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(profileJSON)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func (p *ProfileHandler) Update(w http.ResponseWriter, r *http.Request) {
	//get profileId from Request
	// getProfile -> get profile's techs -> get profile's fields
	// return to user
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

	err = p.pr.Update(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("%+v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(profileJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("%+v\n", err)
		return
	}
	w.Write(b)
}
