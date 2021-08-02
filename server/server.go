package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ueihvn/go-devduo/handler"
	"github.com/ueihvn/go-devduo/service"
)

type Server struct {
	Router         *mux.Router
	profileHandler *handler.ProfileHandler
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
		Router:         mux.NewRouter(),
		profileHandler: handler.NewProfileHandler(repositories.Pr),
	}, nil
}

func (server *Server) Route() {
	server.Router.HandleFunc("/api/v1/user", index).Methods("GET")
	server.Router.HandleFunc("/api/v1/profile", server.profileHandler.Create).Methods("POST")
}
