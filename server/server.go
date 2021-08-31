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
	technologyHandler         *handler.TechnologyHandler
	fieldHandler              *handler.FieldHandler
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
		technologyHandler:         handler.NewTechnologyHandler(repositories.Tr),
		fieldHandler:              handler.NewFieldHandler(repositories.Fr),
	}, nil
}

func (server *Server) Route() {
	//auth
	server.Router.HandleFunc("/signup", server.authHandler.SignUp).Methods("POST")
	server.Router.HandleFunc("/login", server.authHandler.LogIn).Methods("POST")
	server.Router.HandleFunc("/checkauth", server.authHandler.CheckAuthenticate).Methods("GET")
	server.Router.HandleFunc("/refresh", server.authHandler.RefreshCookie).Methods("GET")
	server.Router.HandleFunc("/ok", checkOk).Methods("GET")

	userRouter := server.Router.PathPrefix("/api/v1/user").Subrouter()
	userRouter.HandleFunc("", server.userHandler.Create).Methods("POST")
	userRouter.HandleFunc("", server.userHandler.Update).Methods("PUT")
	userRouter.HandleFunc("/{id:[0-9]+}", server.userHandler.Get).Methods("GET")
	userRouter.Use(server.authHandler.AuthenticateMiddleware, server.authHandler.CheckContentTypeMiddleware, server.authHandler.UserAuthorizeMiddleware)

	profileRouter := server.Router.PathPrefix("/api/v1/profile").Subrouter()
	profileRouter.HandleFunc("", server.profileHandler.Create).Methods("POST")
	profileRouter.HandleFunc("", server.profileHandler.Update).Methods("PUT")
	profileRouter.HandleFunc("/{id:[0-9]+}", server.profileHandler.Create).Methods("GET")
	profileRouter.Use(server.authHandler.AuthenticateMiddleware, server.authHandler.CheckContentTypeMiddleware, server.authHandler.ProfileAuthorizeMiddleware)

	planServiceRouter := server.Router.PathPrefix("/api/v1/planservice").Subrouter()
	planServiceRouter.HandleFunc("", server.planServiceHandler.Create).Methods("POST")
	planServiceRouter.HandleFunc("", server.planServiceHandler.Update).Methods("PUT")
	planServiceRouter.HandleFunc("/{id:[0-9]+}", server.planServiceHandler.Get).Methods("GET")
	planServiceRouter.PathPrefix("").Queries("user_id", "{user_id:[0-9]+}").HandlerFunc(server.planServiceHandler.GetByUserID).Methods("GET")
	profileRouter.Use(server.authHandler.AuthenticateMiddleware, server.authHandler.CheckContentTypeMiddleware, server.authHandler.PlanServiceAuthorizeMiddleware)

	bookingPlanServiceRouter := server.Router.PathPrefix("/api/v1/bookingplanservice").Subrouter()
	bookingPlanServiceRouter.HandleFunc("", server.bookingPlanServiceHandler.Create).Methods("POST")
	bookingPlanServiceRouter.HandleFunc("", server.bookingPlanServiceHandler.Update).Methods("PUT")
	bookingPlanServiceRouter.HandleFunc("/{id:[0-9]+}", server.bookingPlanServiceHandler.Get).Methods("GET")
	bookingPlanServiceRouter.Use(server.authHandler.AuthenticateMiddleware, server.authHandler.CheckContentTypeMiddleware, server.authHandler.BookingPlanServiceAuthorizeMiddleware)

	technologyRouter := server.Router.PathPrefix("/api/v1/technology").Subrouter()
	technologyRouter.HandleFunc("", server.technologyHandler.GetAll).Methods("GET")
	technologyRouter.Use(server.authHandler.AuthenticateMiddleware)

	fieldRouter := server.Router.PathPrefix("/api/v1/field").Subrouter()
	fieldRouter.HandleFunc("", server.fieldHandler.GetAll).Methods("GET")
	fieldRouter.Use(server.authHandler.AuthenticateMiddleware)

	mentorRouter := server.Router.PathPrefix("/api/v1/mentors").Subrouter()
	mentorRouter.HandleFunc("", server.mentorHandler.GetMentors).Methods("GET")
	mentorRouter.Use(server.authHandler.AuthenticateMiddleware)

}
