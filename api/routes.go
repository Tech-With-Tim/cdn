package api

import (
	"net/http"

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
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/docs", http.StatusPermanentRedirect)
		})
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

// Method Route - Handler Function Name
var Routes map[string]string = map[string]string{
	"GET /{AssetUrl}":        "Get Asset",
	"GET /manage/url/{path}": "Fetch Asset Details By URL",
	"GET /manage/id/{id}":    "Fetch Asset Details By ID",
	"POST /manage": "Create Asset",
}
