package handler

import (
	"github.com/akubi0w1/golang-sample/domain/service"
	"github.com/akubi0w1/golang-sample/interface/hash"
	"github.com/akubi0w1/golang-sample/interface/jwt"
	"github.com/akubi0w1/golang-sample/interface/middleware"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	"github.com/akubi0w1/golang-sample/usecase"
	"github.com/go-chi/chi/v5"
)

type App struct {
	// sessionManager session.SessionManager
	user User
}

func NewApp(db *ent.Client) *App {
	userRepository := mysql.NewUser(db)
	hashRepository := hash.NewHash()
	jwtRepository := jwt.NewJWTImpl()

	userService := service.NewUser(userRepository, hashRepository)
	tokenService := service.NewTokenManager(jwtRepository)

	userUsecase := usecase.NewUser(userService, tokenService)

	return &App{
		// sessionManager: session.NewSessionManager(),
		user: NewUser(userUsecase),
	}
}

func (a *App) Routing() *chi.Mux {
	mux := chi.NewRouter()

	mux.Use(middleware.AccessLog)

	// mux.Mount("/", a.authRouter())
	// mux.Mount("/users", a.userRouter())
	mux.Post("/login", a.user.Authorize)

	mux.Route("/users", func(r chi.Router) {
		r.Post("/", a.user.Create)

		r.Group(func(sub chi.Router) {
			sub.Use(middleware.SaveSessionToContext)

			sub.Get("/", a.user.GetAll)
			sub.Get("/users/{userID}", a.user.GetByID)
		})
	})

	mux.Route("/account", func(r chi.Router) {
		r.Group(func(sub chi.Router) {
			sub.Use(middleware.Authorize)

			sub.Put("/", a.user.UpdateProfile)
			sub.Delete("/", a.user.Delete)
		})
	})

	return mux
}

// func (a *App) authRouter() http.Handler {
// 	mux := chi.NewRouter()

// 	mux.Post("/login", a.user.Authorize)

// 	return mux
// }

// func (a *App) userRouter() http.Handler {
// 	mux := chi.NewRouter()

// 	// no session
// 	mux.Post("/users", a.user.Create)

// 	// no auth
// 	mux.Group(func(r chi.Router) {
// 		r.Use(middleware.SaveSessionToContext)

// 		r.Get("/users", a.user.GetAll)
// 		r.Get("/users/{userID}", a.user.GetByID)
// 	})

// 	// required auth
// 	mux.Group(func(r chi.Router) {
// 		r.Use(middleware.Authorize)

// 		r.Put("/account", a.user.UpdateProfile)
// 		r.Delete("/account", a.user.Delete)
// 	})

// 	return mux
// }
