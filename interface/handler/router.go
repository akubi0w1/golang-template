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
	post Post
}

func NewApp(db *ent.Client) *App {
	userRepository := mysql.NewUser(db)
	hashRepository := hash.NewHash()
	jwtRepository := jwt.NewJWTImpl()
	postRepository := mysql.NewPost(db)
	tagRepository := mysql.NewTag(db)
	imageRepository := mysql.NewImage(db)

	userService := service.NewUser(userRepository, hashRepository)
	tokenService := service.NewTokenManager(jwtRepository)
	postService := service.NewPost(postRepository)
	tagService := service.NewTag(tagRepository)
	imageService := service.NewImage(imageRepository)

	userUsecase := usecase.NewUser(userService, tokenService)
	postUsecase := usecase.NewPost(postService, userService, tagService, imageService)

	return &App{
		// sessionManager: session.NewSessionManager(),
		user: NewUser(userUsecase),
		post: NewPost(postUsecase),
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
			sub.Get("/{userID}", a.user.GetByID)
		})
	})

	mux.Route("/account", func(r chi.Router) {
		r.Group(func(sub chi.Router) {
			sub.Use(middleware.Authorize)

			sub.Put("/", a.user.UpdateProfile)
			sub.Delete("/", a.user.Delete)
		})
	})

	mux.Route("/posts", func(r chi.Router) {
		r.Get("/", a.post.GetAll)
		r.Get("/{postID}", a.post.GetByID)

		r.Group(func(sub chi.Router) {
			sub.Use(middleware.Authorize)

			sub.Post("/", a.post.Create)
		})
	})

	return mux
}
