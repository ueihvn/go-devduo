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
	authHandler               *handler.AuthHandler
	mentorHandler             *handler.MentorHandler
}

func checkOk(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func NewServer() (*Server, error) {
	repositories, err := service.NewRepositories()
	if err != nil {
		return nil, err
	}

	err = repositories.InitData()
	if err != nil {
		return nil, err
	}

	return &Server{
		Router:                    mux.NewRouter(),
		profileHandler:            handler.NewProfileHandler(repositories.Pr, repositories.Tr, repositories.Fr),
		userHandler:               handler.NewUserHandler(repositories.Ur),
		planServiceHandler:        handler.NewPlanServiceHandler(repositories.Psr),
		bookingPlanServiceHandler: handler.NewBookingPlanServiceHandler(repositories.Bpsr),
		authHandler:               handler.NewAuthHandler(repositories.Ur),
		mentorHandler:             handler.NewMentorHandler(repositories.Bpsr, repositories.Psr, repositories.Pr),
	}, nil
}

func (server *Server) Route() {
	//auth
	server.Router.HandleFunc("/signup", server.authHandler.SignUp).Methods("POST")
	server.Router.HandleFunc("/login", server.authHandler.LogIn).Methods("POST")
	server.Router.HandleFunc("/checkauth", server.authHandler.CheckAuthenticate).Methods("GET")
	server.Router.HandleFunc("/ok", checkOk).Methods("GET")

	// + subrouter for api
	subRouter := server.Router.PathPrefix("/api/v1").Subrouter()
	//user
	subRouter.HandleFunc("/user", server.userHandler.Create).Methods("POST")
	subRouter.HandleFunc("/user/{id:[0-9]+}", server.userHandler.Get).Methods("GET")
	// add authorization
	subRouter.HandleFunc("/user", server.userHandler.Update).Methods("PUT")

	//profile
	// add authorization
	subRouter.HandleFunc("/profile", server.profileHandler.Create).Methods("POST")
	subRouter.HandleFunc("/profile", server.profileHandler.Update).Methods("PUT")

	subRouter.HandleFunc("/profile/{id:[0-9]+}", server.profileHandler.Get).Methods("GET")

	//planservice
	// add authorization
	subRouter.HandleFunc("/planservice", server.planServiceHandler.Create).Methods("POST")
	subRouter.HandleFunc("/planservice", server.planServiceHandler.Update).Methods("PUT")

	subRouter.HandleFunc("/planservice/{id:[0-9]+}", server.planServiceHandler.Get).Methods("GET")
	subRouter.HandleFunc("/planservice/{user_id:[0-9]+}", server.planServiceHandler.GetByUserID).Methods("GET")

	//bookingplanservice
	// add authorization
	subRouter.HandleFunc("/bookingplanservice", server.bookingPlanServiceHandler.Create).Methods("POST")
	subRouter.HandleFunc("/bookingplanservice", server.bookingPlanServiceHandler.Update).Methods("PUT")

	subRouter.HandleFunc("/bookingplanservice/{id:[0-9]+}", server.bookingPlanServiceHandler.Get).Methods("GET")
	// +bookingplanservice by PlanServiceId || menteeid

	// mentor
	subRouter.PathPrefix("/mentors").HandlerFunc(server.mentorHandler.GetMentors).Methods("GET")

	subRouter.Use(server.authHandler.AuthenticateMiddleware)

}
