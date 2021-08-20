package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

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

	// + validate
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
	userId, err := parseID(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

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

func (u *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	err = u.ur.UpdateUser(&user)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	responseJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Printf("%+v\n", err)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func (u *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user model.User
	err := FromJSON(&user, r.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "err deserialize data. Check request"}, w)
		return
	}

	// + validate user

	// + hash password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: "err hash password"}, w)
		return
	}
	user.Password = hashedPassword

	// create database
	err = u.ur.CreateUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			err = errors.New("already exits with email")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err = errors.New("database error")
			w.WriteHeader(http.StatusInternalServerError)
		}
		ToJSON(Response{Status: false, Message: err.Error()}, w)
		return

	}
	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully create user", Data: IdResponse{Id: user.ID}}, w)
}

func (u *UserHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user model.User
	err := FromJSON(&user, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "err deserialize data. Check request"}, w)
		return
	}

	// + validate user
	userData, err := u.ur.GetUserByEmail(user.Email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = errors.New("user not found with email")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err = errors.New("database error")
			w.WriteHeader(http.StatusInternalServerError)
		}

		ToJSON(Response{Status: false, Message: err.Error()}, w)
		return
	}
	// check hash
	err = ComparePassword(userData.Password, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "wrong password with email"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully authorization user", Data: userData}, w)
}
