package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ueihvn/go-devduo/handler"
	"github.com/ueihvn/go-devduo/service"
)

type Server struct {
	Router                    *mux.Router
	profileHandler            *handler.ProfileHandler
	userHandler               *handler.UserHandler
	planServiceHandler        *handler.PlanServiceHandler
	bookingPlanServiceHandler *handler.BookingPlanServiceHandler
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func NewServer() (*Server, error) {
	repositories, err := service.NewRepositories()
	if err != nil {
		return nil, err
	}

	return &Server{
		Router:                    mux.NewRouter(),
		profileHandler:            handler.NewProfileHandler(repositories.Pr, repositories.Tr, repositories.Fr),
		userHandler:               handler.NewUserHandler(repositories.Ur),
		planServiceHandler:        handler.NewPlanServiceHandler(repositories.Psr),
		bookingPlanServiceHandler: handler.NewBookingPlanServiceHandler(repositories.Bpsr),
	}, nil
}

func (server *Server) Route() {
	server.Router.HandleFunc("/api/v1/ok", index).Methods("GET")

	//user
	server.Router.HandleFunc("/api/v1/user", server.userHandler.Create).Methods("POST")
	server.Router.HandleFunc("/api/v1/user/{id:[0-9]+}", server.userHandler.Get).Methods("GET")
	server.Router.HandleFunc("/api/v1/user", server.userHandler.Update).Methods("PUT")

	//profile
	server.Router.HandleFunc("/api/v1/profile/{id:[0-9]+}", server.profileHandler.Get).Methods("GET")
	server.Router.HandleFunc("/api/v1/profile", server.profileHandler.Create).Methods("POST")
	server.Router.HandleFunc("/api/v1/profile", server.profileHandler.Update).Methods("PUT")

	//planservice
	server.Router.HandleFunc("/api/v1/planservice", server.planServiceHandler.Create).Methods("POST")
	server.Router.HandleFunc("/api/v1/planservice/{id:[0-9]+}", server.planServiceHandler.Get).Methods("GET")
	server.Router.HandleFunc("/api/v1/planservice", server.planServiceHandler.Update).Methods("PUT")

	//bookingplanservice
	server.Router.HandleFunc("/api/v1/bookingplanservice", server.bookingPlanServiceHandler.Create).Methods("POST")
	server.Router.HandleFunc("/api/v1/bookingplanservice/{id:[0-9]+}", server.bookingPlanServiceHandler.Get).Methods("GET")
	server.Router.HandleFunc("/api/v1/bookingplanservice", server.bookingPlanServiceHandler.Update).Methods("PUT")

}
