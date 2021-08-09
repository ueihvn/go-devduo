package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ueihvn/go-devduo/model"
)

type UserHandler struct {
	ur model.UserRepository
}

func NewUserHandler(userRepository model.UserRepository) *UserHandler {
	return &UserHandler{
		ur: userRepository,
	}
}

// create handler function
func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	b := r.ContentLength
	body := make([]byte, b)
	r.Body.Read(body)

	var user model.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = u.ur.CreateUser(&user)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseMap := map[string]interface{}{
		"id": user.ID,
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

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := parseID(vars["id"])

	user, err := u.ur.GetUserById(userId)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
