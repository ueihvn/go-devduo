package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

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

func (ah *AuthHandler) UserAuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Methods") != "PUT" {
			next.ServeHTTP(w, r)
			return
		}

		var user model.User
		err := FromJSON(&user, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: parseJsonError}, w)
			return
		}
		session, _ := ah.Store.Get(r, "session.id")
		if session.Values["user_id"] != user.ID {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: "Unauthorized,can not update other user"}, w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (ah *AuthHandler) ProfileAuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if methods := r.Header.Get("Methods"); methods != "PUT" && methods != "POST" {
			next.ServeHTTP(w, r)
			return
		}

		var profile model.ProfileJSON
		err := FromJSON(&profile, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: parseJsonError}, w)
			return
		}
		session, _ := ah.Store.Get(r, "session.id")
		if session.Values["user_id"] != profile.UserID {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: "Unauthorized,can not create or update other profile"}, w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (ah *AuthHandler) PlanServiceAuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if methods := r.Header.Get("Methods"); methods != "PUT" && methods != "POST" {
			next.ServeHTTP(w, r)
			return
		}

		var ps model.PlanService
		err := FromJSON(&ps, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: parseJsonError}, w)
			return
		}
		session, _ := ah.Store.Get(r, "session.id")
		if session.Values["user_id"] != ps.UserID {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: "Unauthorized,can not create or update other user PlanService"}, w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (ah *AuthHandler) BookingPlanServiceAuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if methods := r.Header.Get("Methods"); methods != "PUT" && methods != "POST" {
			next.ServeHTTP(w, r)
			return
		}

		var bps model.BookingPlanService
		err := FromJSON(&bps, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: parseJsonError}, w)
			return
		}
		session, _ := ah.Store.Get(r, "session.id")
		if session.Values["user_id"] != bps.UserID {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: "Unauthorized,can not create or update other user booking plan service"}, w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (ah *AuthHandler) CheckContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if methods := r.Header.Get("Methods"); methods != "PUT" && methods != "POST" {
			next.ServeHTTP(w, r)
			return
		}

		if contentType := r.Header.Get("Content-type"); contentType != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			ToJSON(Response{Status: false, Message: "content-type must be application/json"}, w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// swagger:route POST /signup auth signupLoginReq
// signup with email and password
// responses:
// 200: signUpResp
// 400: errorResp
// 500: errorResp
func (ah *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user model.User
	err := FromJSON(&user, r.Body)
	if err != nil {
		fmt.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
		return
	}

	// + validate user

	// + hash password
	hashedPassword, err := user.HashPassword()
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

// swagger:route POST /login auth signupLoginReq
// login with email and password
// responses:
// 200: signUpResp
// 400: errorResp
// 500: errorResp
func (ah *AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var user model.User
	err := FromJSON(&user, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
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
	err = userData.ComparePassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: "wrong password with email"}, w)
		return
	}

	// return cookie
	session, _ := ah.Store.Get(r, "session.id")
	session.Values["authenticated"] = true
	const SESSION_MAXAGE = 300

	session.Values["user_id"] = userData.ID
	session.Options.MaxAge = SESSION_MAXAGE
	session.Options.HttpOnly = true
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
		fmt.Printf("%+v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully authorization user", Data: IdResponse{Id: userData.ID}}, w)
}

// swagger:route GET /refresh auth noReq
// refresh cookie
// responses:
// 200: refreshCookieResp
// 400: errorResp
// 401: errorResp
// 500: errorResp
func (ah *AuthHandler) RefreshCookie(w http.ResponseWriter, r *http.Request) {
	session, _ := ah.Store.Get(r, "session.id")
	if session.Values["authenticated"] != nil && session.Values["authenticated"] != false {
		const SESSION_MAXAGE = 300
		session.Options.MaxAge = SESSION_MAXAGE
		session.Options.HttpOnly = true
		err := session.Save(r, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ToJSON(Response{Status: false, Message: err.Error()}, w)
			fmt.Printf("%+v\n", err)
			return
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		ToJSON(Response{Status: true, Message: "successfully refresh cookie"}, w)
	} else {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		ToJSON(Response{Status: false, Message: "Unauthorized, please login"}, w)
	}
}
