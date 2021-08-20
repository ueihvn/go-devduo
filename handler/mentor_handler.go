package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
	"github.com/ueihvn/go-devduo/model"
)

type MentorHandler struct {
	bpsr model.BookingPlanServiceRepository
	psr  model.PlanServiceRepository
	pr   model.ProfileRepository
}

type Mentor struct {
	model.ProfileJSON
	Price  decimal.Decimal `json:"price"`
	Mentee uint64          `json:"mentee"`
}

func NewMentorHandler(bpsRepo model.BookingPlanServiceRepository, psRepo model.PlanServiceRepository, pRepo model.ProfileRepository) *MentorHandler {
	return &MentorHandler{
		bpsr: bpsRepo,
		psr:  psRepo,
		pr:   pRepo,
	}
}

func (mh *MentorHandler) getMentor(profile *model.Profile) (*Mentor, error) {
	var mentor Mentor
	profileJSON, err := profile.ToProfileJSON()
	if err != nil {
		return nil, err
	}

	mentor.ProfileJSON = *profileJSON

	price, err := mh.psr.GetSmallestPricePlanServiceByUserID(mentor.UserID)
	if err != nil {
		return nil, err
	}
	mentor.Price = *price

	mentee, err := mh.bpsr.CountUserBookPlanServiceByUserID(mentor.UserID)
	if err != nil {
		return nil, err
	}
	mentor.Mentee = mentee

	return &mentor, nil

}

func (mh *MentorHandler) Get(w http.ResponseWriter, r *http.Request) {
	// get mentor from offset to limit
	w.Header().Set("Content-type", "application/json")
	query := mux.Vars(r)
	offset, _ := parseID(query["o"])
	limit, _ := parseID(query["l"])

	var mentors []Mentor
	profiles, err := mh.pr.GetFromOffsetToLimitOfProfile(int(offset), int(limit))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	for _, profile := range profiles {
		mentor, err := mh.getMentor(&profile)
		if err != nil {

		}

		mentors = append(mentors, *mentor)
	}

	ToJSON(Response{
		Status:  true,
		Message: "mentor from offset to limit",
		Data:    mentors,
	}, w)

}

func (mh *MentorHandler) GetWithLimit(w http.ResponseWriter, r *http.Request) {
	// get mentor from offset to limit
	w.Header().Set("Content-type", "application/json")
	query := mux.Vars(r)
	limit, err := parseID(query["l"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	var mentors []Mentor
	profiles, err := mh.pr.GetFromOffsetToLimitOfProfile(0, int(limit))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	for _, profile := range profiles {
		mentor, err := mh.getMentor(&profile)
		if err != nil {

		}

		mentors = append(mentors, *mentor)
	}

	ToJSON(Response{
		Status:  true,
		Message: "mentor from offset to limit",
		Data:    mentors,
	}, w)

}

func (mh *MentorHandler) GetWithLimitLastID(w http.ResponseWriter, r *http.Request) {
	// get mentor from offset to limit
	w.Header().Set("Content-type", "application/json")
	query := mux.Vars(r)
	lastID, _ := parseID(query["last_id"])
	limit, _ := parseID(query["l"])

	var mentors []Mentor
	profiles, err := mh.pr.GetWithLimitLastID(int(limit), lastID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	for _, profile := range profiles {
		mentor, err := mh.getMentor(&profile)
		if err != nil {

		}

		mentors = append(mentors, *mentor)
	}

	ToJSON(Response{
		Status:  true,
		Message: "mentor from offset to limit",
		Data:    mentors,
	}, w)

}
