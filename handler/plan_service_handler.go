package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ueihvn/go-devduo/model"
)

type PlanServiceHandler struct {
	psr model.PlanServiceRepository
}

func NewPlanServiceHandler(psRepo model.PlanServiceRepository) *PlanServiceHandler {
	return &PlanServiceHandler{
		psr: psRepo,
	}
}

func (psH *PlanServiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	b := r.ContentLength
	body := make([]byte, b)
	r.Body.Read(body)

	var ps model.PlanService
	err := json.Unmarshal(body, &ps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// + validate
	err = psH.psr.Create(&ps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	responseMap := map[string]interface{}{
		"id": ps.ID,
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

func (psH *PlanServiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	psID, err := parseID(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	planService, err := psH.psr.Get(psID)
	psJSON, err := json.Marshal(planService)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(psJSON)
}

func (psH *PlanServiceHandler) Update(w http.ResponseWriter, r *http.Request) {
	b := r.ContentLength
	body := make([]byte, b)
	r.Body.Read(body)

	var ps model.PlanService
	err := json.Unmarshal(body, &ps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = psH.psr.Update(&ps)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseJSON, err := json.Marshal(ps)
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
