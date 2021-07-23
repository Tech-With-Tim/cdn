package api

import (
	"github.com/Tech-With-Tim/cdn/api/handlers"
	"github.com/Tech-With-Tim/cdn/cache"
	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
)

var postCache cache.PostCache

func MainRouter(r *chi.Mux, store *db.Store, config utils.Config) {

	postCache = cache.NewRedisCache(
		config.RedisHost,
		config.RedisDb,
		config.RedisPass,
		60)

	r.Group(func(r chi.Router) {
		r.Get("/", handlers.HelloWorld())
		r.Get("/{AssetUrl}", handlers.GetAsset(store, postCache))
		r.Get("/manage/url/{path}", handlers.FetchAssetDetailsByURL(store))
		r.Get("/manage/id/{path}", handlers.FetchAssetDetailsByID(store))
		r.Get("/docs", handlers.GetDocs())
	})

	//Private Routes
	r.Group(func(r chi.Router) {
		r.Use(AuthJwtWrap(config.SecretKey))
		r.Post("/manage", handlers.CreateAsset(store, config.MaxFileSize))
		r.Get("/testing", handlers.HelloWorld())
	})

}
