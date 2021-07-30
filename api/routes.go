package api

import (
	"github.com/Tech-With-Tim/cdn/api/handlers"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
)

// var postCache cache.PostCache

func MainRouter(r *chi.Mux, config utils.Config, s handlers.Handler) {

	// postCache = cache.NewRedisCache(
	// 	config.RedisHost,
	// 	config.RedisDb,
	// 	config.RedisPass,
	// 	60)

	r.Group(func(r chi.Router) {
		r.Get("/docs", handlers.GetDocs())
		
		r.Get("/{StaticFile}", handlers.ServeDocsPage())
		r.Get("/build/{StaticFile}", handlers.ServeDocsPage())
		r.Get("/assets/{StaticFile}", handlers.ServeDocsPage())
		r.Get("/", handlers.ServeDocsPage())

		r.Get("/{AssetUrl}", s.GetAsset())
		r.Get("/manage/url/{path}", s.FetchAssetDetailsByURL())
		r.Get("/manage/id/{id}", s.FetchAssetDetailsByID())
	})

	//Private Routes
	r.Group(func(r chi.Router) {
		r.Use(AuthJwtWrap(config.SecretKey))
		r.Post("/manage", s.CreateAsset(config.MaxFileSize))
		r.Get("/testing", handlers.HelloWorld())
	})
}
