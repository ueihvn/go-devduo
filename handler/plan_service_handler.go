package handler

import (
	"net/http"
	"strings"

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

// swagger:route POST /api/v1/planservice planservice planServiceCreateUpdateReq
// create plan service
// responses:
// 200: planServiceCreateUpdateResp
// 400: errorResp
// 500: errorResp
func (psH *PlanServiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ps model.PlanService
	err := FromJSON(&ps, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
		return
	}

	// + validate

	err = psH.psr.Create(&ps)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	ToJSON(Response{Status: true, Message: "successfully create plan service", Data: ps}, w)
}

// swagger:operation GET /api/v1/planservice/{id} planservice get
// ---
// summary: Get planservice by planservice id
// parameters:
// - name: id
//   in: path
//   description: id of planservice
//   required: true
//   type: integer
//   format: unit64
// responses:
//   "200":
//     "$ref": "#/responses/planServiceCreateUpdateResp"
//   "403":
//     "$ref": "#/responses/errorResp"
//   "500":
//     "$ref": "#/responses/errorResp"
func (psH *PlanServiceHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	psID, err := parseID(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: ParseIDError}, w)
		return
	}

	planService, err := psH.psr.Get(psID)
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
	ToJSON(Response{Status: true, Message: "successfully get plan serivce", Data: planService}, w)
}

// swagger:route PUT /api/v1/planservice planservice planServiceCreateUpdateReq
// update planservice
// responses:
// 200: planServiceCreateUpdateResp
// 400: errorResp
// 500: errorResp
func (psH *PlanServiceHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ps model.PlanService
	err := FromJSON(&ps, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: parseJsonError}, w)
		return
	}

	err = psH.psr.Update(&ps)
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
	ToJSON(Response{Status: true, Message: "successfully update plan service", Data: ps}, w)
}

// swagger:operation GET /api/v1/planservice planservice getPlanServiceOfUser
// ---
// summary: Get all planservice of user
// parameters:
// - name: user_id
//   in: query
//   description: user id
//   required: true
//   type: integer
//   format: unit64
// responses:
//   "200":
//     "$ref": "#/responses/planServiceGetListResp"
//   "403":
//     "$ref": "#/responses/errorResp"
//   "500":
//     "$ref": "#/responses/errorResp"
func (psH *PlanServiceHandler) GetPlanServiceOfUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userID, err := parseID(r.URL.Query().Get("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ToJSON(Response{Status: false, Message: ParseIDError}, w)
	}

	planServices, err := psH.psr.GetPlanServiceByUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ToJSON(Response{Status: false, Message: err.Error()}, w)
	}

	ToJSON(Response{
		Status:  true,
		Message: "get planservice by userID",
		Data:    planServices,
	}, w)

}
