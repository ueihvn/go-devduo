package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ueihvn/go-devduo/model"
)

type ProfileHandler struct {
	fr model.ProfileRepository
}

func NewProfileHandler(profileRepository model.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{
		fr: profileRepository,
	}
}

func (p *ProfileHandler) Create(w http.ResponseWriter, r *http.Request) {
	var profile model.Profile
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &profile)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(profile)
	}

	err = p.fr.Create(profile)
	if err != nil {
		fmt.Println(err)
	}
}
