package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/ueihvn/go-devduo/model"
)

type AuthHandler struct {
	Store *sessions.CookieStore
	ur    model.UserRepository
}

func NewAuthHandler(uRepo model.UserRepository) *AuthHandler {
	return &AuthHandler{
		ur:    uRepo,
		Store: sessions.NewCookieStore([]byte("secret_key")),
	}
}

func (ah *AuthHandler) AuthenticateMiddleware(next http.Handler) http.Handler {
	//before running
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := ah.Store.Get(r, "session.id")
		if session.Values["authenticated"] != nil && session.Values["authenticated"] != false {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			ToJSON(Response{Status: false, Message: "Unauthorized, please login"}, w)
		}
	})
}

func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
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
	err = ah.ur.CreateUser(&user)
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

func (ah *AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user model.User
	err := FromJSON(&user, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "err deserialize data. Check request"}, w)
		return
	}

	// + validate user
	userData, err := ah.ur.GetUserByEmail(user.Email)
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

	// return cookie
	session, err := ah.Store.Get(r, "session.id")
	if err != nil {
		fmt.Println(err)
	}

	session.Values["authenticated"] = true
	session.Options.MaxAge = 60
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
		fmt.Printf("%+v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully authorization user", Data: userData}, w)
}

func (ah *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	session, _ := ah.Store.Get(r, "session.id")
	if session.Values["authenticated"] != nil && session.Values["authenticated"] != false {
		fmt.Println(session)
		w.Write([]byte(time.Now().String()))
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ToJSON(Response{Status: false, Message: "Unauthorized, please login"}, w)
	}
}
