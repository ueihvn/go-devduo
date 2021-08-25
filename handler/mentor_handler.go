package handler

import (
	"fmt"
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

func (mh *MentorHandler) GetWithLimitOffset(w http.ResponseWriter, r *http.Request) {
	// get mentor from offset to limit
	w.Header().Set("Content-type", "application/json")
	query := mux.Vars(r)
	offset, _ := parseID(query["o"])
	limit, _ := parseID(query["l"])

	fmt.Println("GetWithLimitOffset")

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

	fmt.Println("GetWithLimit")

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

func (mh *MentorHandler) GetWithLimitCursor(w http.ResponseWriter, r *http.Request) {
	// get mentor from offset to limit
	w.Header().Set("Content-type", "application/json")
	query := mux.Vars(r)
	cursor, _ := parseID(query["cursor"])
	limit, _ := parseID(query["l"])

	fmt.Println("GetWithLimitCursor")

	var mentors []Mentor
	profiles, err := mh.pr.GetWithLimitLastID(int(limit), cursor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	for _, profile := range profiles {
		mentor, err := mh.getMentor(&profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ToJSON(Response{Status: false, Message: err.Error()}, w)
		}

		mentors = append(mentors, *mentor)
	}

	ToJSON(Response{
		Status:  true,
		Message: "mentor from offset to limit",
		Data:    mentors,
	}, w)

}

func (mh *MentorHandler) FilterMentorsByFields(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	test := r.URL.Query()
	fields, err := fromStrIDsToArrUnitIDs(test["field"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "faild to parse field_id. Chech request"}, w)
	}

	profiles, err := mh.pr.FilterProfileByFields(fields)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	var mentors []Mentor

	for _, profile := range profiles {
		mentor, err := mh.getMentor(&profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ToJSON(Response{Status: false, Message: err.Error()}, w)
		}

		mentors = append(mentors, *mentor)
	}

	ToJSON(Response{
		Status:  true,
		Message: "mentor with field",
		Data:    mentors,
	}, w)

}

func (mh *MentorHandler) FilterMentorsByTechs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	tech, err := fromStrIDsToArrUnitIDs(r.URL.Query().Get("tech"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "faild to parse tech_id. Chech request"}, w)
	}

	profiles, err := mh.pr.FilterProfileByTechs(tech)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	var mentors []Mentor

	for _, profile := range profiles {
		mentor, err := mh.getMentor(&profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ToJSON(Response{Status: false, Message: err.Error()}, w)
		}

		mentors = append(mentors, *mentor)
	}

	ToJSON(Response{
		Status:  true,
		Message: "mentor with tech",
		Data:    mentors,
	}, w)

}

func (mh *MentorHandler) FilterMentorsByFieldsTechs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	test := r.URL.Query()
	fields, err := fromStrIDsToArrUnitIDs(test["field"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "faild to parse field_id. Chech request"}, w)
	}

	techs, err := fromStrIDsToArrUnitIDs(test["tech"][0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "faild to parse tech_id. Chech request"}, w)
	}

	profiles, err := mh.pr.FilterProfileByFieldsTechs(fields, techs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	var mentors []Mentor

	for _, profile := range profiles {
		mentor, err := mh.getMentor(&profile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ToJSON(Response{Status: false, Message: err.Error()}, w)
		}

		mentors = append(mentors, *mentor)
	}

	ToJSON(Response{
		Status:  true,
		Message: "mentor with field&tech",
		Data:    mentors,
	}, w)

}

func (mh *MentorHandler) GetMentors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	strTechs := r.URL.Query().Get("tech")
	if strTechs != "" {
		fmt.Println(strTechs)
	}
	strFields := r.URL.Query().Get("field")
	if strFields != "" {
		fmt.Println(strFields)
	}

}
