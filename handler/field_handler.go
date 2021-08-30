package handler

import (
	"net/http"

	"github.com/ueihvn/go-devduo/model"
)

type FieldHandler struct {
	fr model.FieldRepository
}

func NewFieldHandler(fRepo model.FieldRepository) *FieldHandler {
	return &FieldHandler{
		fr: fRepo,
	}
}

func (fh *FieldHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	fields, err := fh.fr.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "err deserialize data. Check request"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully get all technologies", Data: fields}, w)
}
