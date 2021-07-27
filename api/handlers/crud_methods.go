package handlers

import (
	"context"

	db "github.com/Tech-With-Tim/cdn/db/sqlc"
)

func (s *stores) getFile(url string,
	ctx context.Context) (fileRow db.GetFileRow, err error) {

	cachedFile := s.Cache.Get(url)
	if cachedFile == nil {
		fileRow, err = s.Store.GetFile(ctx, url)

		if err != nil {
			return
		}
		s.Cache.Set(url, &fileRow)

	} else {
		fileRow = *cachedFile
	}

	return
}
