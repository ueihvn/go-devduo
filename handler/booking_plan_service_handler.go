package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ueihvn/go-devduo/model"
)

type BookingPlanServiceHandler struct {
	bpsr model.BookingPlanServiceRepository
}

func NewBookingPlanServiceHandler(bpsRepo model.BookingPlanServiceRepository) *BookingPlanServiceHandler {
	return &BookingPlanServiceHandler{
		bpsr: bpsRepo,
	}
}

func (bpsH *BookingPlanServiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	b := r.ContentLength
	body := make([]byte, b)
	r.Body.Read(body)

	var bps model.BookingPlanService
	err := json.Unmarshal(body, &bps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// + validate
	err = bpsH.bpsr.Create(&bps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseMap := map[string]interface{}{
		"id": bps.ID,
	}

	responseJSON, err := json.Marshal(responseMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("%+v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func (bpsH *BookingPlanServiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bpsID, err := parseID(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	bookingPlanService, err := bpsH.bpsr.Get(bpsID)
	bpsJSON, err := json.Marshal(bookingPlanService)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bpsJSON)
}

func (bpsH *BookingPlanServiceHandler) Update(w http.ResponseWriter, r *http.Request) {

	b := r.ContentLength
	body := make([]byte, b)
	r.Body.Read(body)

	var bps model.BookingPlanService
	err := json.Unmarshal(body, &bps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = bpsH.bpsr.Update(&bps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseJSON, err := json.Marshal(bps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("%+v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
