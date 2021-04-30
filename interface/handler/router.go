package handler

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/domain/service"
	"github.com/akubi0w1/golang-sample/interface/middleware"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	"github.com/akubi0w1/golang-sample/interface/session"
	"github.com/akubi0w1/golang-sample/usecase"
	"github.com/go-chi/chi/v5"
)

type App struct {
	sessionManager session.SessionManager
	user           User
}

func NewApp(db *ent.Client) *App {
	userRepository := mysql.NewUser(db)

	userService := service.NewUser(userRepository)

	userUsecase := usecase.NewUser(userService)

	return &App{
		sessionManager: session.NewSessionManager(),
		user:           NewUser(userUsecase),
	}
}

func (a *App) Routing() {
	mux := chi.NewRouter()
	mux.Use(middleware.AccessLog)

	mux.Mount("/", a.authRouter())
	mux.Mount("/users", a.userRouter())
}

func (a *App) authRouter() http.Handler {
	mux := chi.NewRouter()
	return mux
}

func (a *App) userRouter() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", a.user.GetAll)
	return mux
}
