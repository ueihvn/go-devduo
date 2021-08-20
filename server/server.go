package server

import (
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
		authHandler:               handler.NewAuthHandler(repositories.Ur),
		mentorHandler:             handler.NewMentorHandler(repositories.Bpsr, repositories.Psr, repositories.Pr),
	}, nil
}

func (server *Server) Route() {
	//auth
	server.Router.HandleFunc("/signup", server.authHandler.SignUp).Methods("POST")
	server.Router.HandleFunc("/login", server.authHandler.LogIn).Methods("POST")
	server.Router.HandleFunc("/checkauth", server.authHandler.CheckAuth).Methods("GET")

	// + subrouter for api
	subRouter := server.Router.PathPrefix("/api/v1").Subrouter()
	//user
	subRouter.HandleFunc("/user", server.userHandler.Create).Methods("POST")
	subRouter.HandleFunc("/user/{id:[0-9]+}", server.userHandler.Get).Methods("GET")
	subRouter.HandleFunc("/user", server.userHandler.Update).Methods("PUT")

	//profile
	subRouter.HandleFunc("/profile/{id:[0-9]+}", server.profileHandler.Get).Methods("GET")
	subRouter.HandleFunc("/profile", server.profileHandler.Create).Methods("POST")
	subRouter.HandleFunc("/profile", server.profileHandler.Update).Methods("PUT")

	//planservice
	subRouter.HandleFunc("/planservice", server.planServiceHandler.Create).Methods("POST")
	subRouter.HandleFunc("/planservice/{id:[0-9]+}", server.planServiceHandler.Get).Methods("GET")
	subRouter.HandleFunc("/planservice", server.planServiceHandler.Update).Methods("PUT")
	// +planservice by userid

	//bookingplanservice
	subRouter.HandleFunc("/bookingplanservice", server.bookingPlanServiceHandler.Create).Methods("POST")
	subRouter.HandleFunc("/bookingplanservice/{id:[0-9]+}", server.bookingPlanServiceHandler.Get).Methods("GET")
	subRouter.HandleFunc("/bookingplanservice", server.bookingPlanServiceHandler.Update).Methods("PUT")
	// +bookingplanservice by PlanServiceId || mentee

	// subRouter.Use(server.authHandler.AuthenticateMiddleware)

	// mentor
	subRouter.PathPrefix("/mentors").Queries("l", "{l:[0-9]+}").HandlerFunc(server.mentorHandler.GetWithLimit).Methods("GET")
	subRouter.PathPrefix("/mentors").Queries("l", "{l:[0-9]+}", "last_id", "{last_id:[0-9]+}").HandlerFunc(server.mentorHandler.GetWithLimitLastID).Methods("GET")
	subRouter.PathPrefix("/mentors").Queries("o", "{o:[0-9]+}", "l", "{l:[0-9]+}").HandlerFunc(server.mentorHandler.Get).Methods("GET")
}
