package handler

import (
	"net/http"

	"github.com/ueihvn/go-devduo/model"
)

type TechnologyHandler struct {
	tr model.TechnologyRepository
}

func NewTechnologyHandler(tRepo model.TechnologyRepository) *TechnologyHandler {
	return &TechnologyHandler{
		tr: tRepo,
	}
}

func (th *TechnologyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	technologies, err := th.tr.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "err deserialize data. Check request"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully get all technologies", Data: technologies}, w)
}
