package handlers

import (
	"context"
	"net/http"

	"github.com/Tech-With-Tim/cdn/cache"
	db "github.com/Tech-With-Tim/cdn/db/sqlc"
)

type stores struct {
	Store *db.Store
	Cache cache.PostCache
}

type Service struct {
	*stores
}

type DBHandler interface {
	getFile(url string, ctx context.Context) (db.GetFileRow, error)
}

type Handler interface {
	FetchAssetDetailsByURL() http.HandlerFunc
	FetchAssetDetailsByID() http.HandlerFunc
	CreateAsset(FileSize int64) http.HandlerFunc
	GetAsset() http.HandlerFunc
}

func NewServiceHandler(store *db.Store, cache cache.PostCache) Handler {
	return &Service{
		&stores{
			Store: store,
			Cache: cache,
		},
	}
}
