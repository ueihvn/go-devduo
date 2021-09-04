package handler

import (
	"encoding/json"
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

// swagger:operation GET /api/v1/user/{id} user get
// ---
// summary: Get user information by user id
// parameters:
// - name: id
//   in: path
//   description: id of user
//   required: true
//   type: integer
//   format: unit64
// responses:
//   "200":
//     "$ref": "#/responses/userResp"
//   "403":
//     "$ref": "#/responses/errorResp"
//   "500":
//     "$ref": "#/responses/errorResp"
func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userId, err := parseID(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: ParseIDError}, w)
		return
	}

	user, err := u.ur.GetUserById(userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
			ToJSON(Response{Status: false, Message: notFoundError}, w)
		}
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: serverInternalError}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully get user by id", Data: user}, w)
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
