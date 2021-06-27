package cache

import db "github.com/Tech-With-Tim/cdn/db/sqlc"

type PostCache interface {
	Set(key string, value *db.GetFileRow)
	Get(key string) *db.GetFileRow
}
