package api

import (
	"github.com/Ibezio/cdn/api/handlers"
	db "github.com/Ibezio/cdn/db/sqlc"
	"github.com/Ibezio/cdn/utils"
	"github.com/go-chi/chi/v5"
)

func MainRouter(r *chi.Mux, store *db.Store, config utils.Config) {
	//tokenAuth = jwtauth.New("HS256", []byte("bruh"), nil)
	//Declare Handlers Here
	//Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", handlers.HelloWorld())
		r.Get("/{AssetUrl}", handlers.GetAsset(store))
		//r.Get("/manage/{path}", FetchAssetDetails(db))
	})

	//Private Routes
	r.Group(func(r chi.Router) {
		r.Use(AuthJwtWrap(config.SecretKey, config.JwtIssuer))
		r.Post("/manage", handlers.CreateAsset(store, config.MaxFileSize))
		r.Get("/testing", handlers.HelloWorld())
	})
}