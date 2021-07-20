package cache

import db "github.com/Tech-With-Tim/cdn/db/sqlc"

// PostCache Implements Cache Functions
type PostCache interface {
	Set(key string, value *db.GetFileRow)
	Get(key string) *db.GetFileRow
}
