package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Tech-With-Tim/cdn/cache"
	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
)

func GetAsset(store *db.Store, cache cache.PostCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		url := chi.URLParam(r, "AssetUrl")
		var fileRow db.GetFileRow
		var err error
		cachedFile := cache.Get(url)
		if cachedFile == nil {
			fileRow, err = store.GetFile(r.Context(), url)
			if err != nil {
				if err == sql.ErrNoRows {
					resp = map[string]interface{}{"error": "Not Found",
						"message": "No asset found with that url_path."}
					utils.JSON(w, http.StatusNotFound, resp)
					return
				}
				resp = map[string]interface{}{"error": "Something Unexpected Occurred."}
				utils.JSON(w, http.StatusInternalServerError, resp)
				log.Println(err.Error())
				return
			}
			cache.Set(url, &fileRow)
		} else {
			//fmt.Println("Found File in Redis Cache 🍻")
			fileRow = *cachedFile
		}

		// FileRow:
		//    Data     []byte `json:"data"`
		//    Mimetype string `json:"mimetype"`

		w.Header().Set("Content-Type", fileRow.Mimetype)
		_, err = w.Write(fileRow.Data)
		if err != nil {
			resp = map[string]interface{}{"error": "Something Unexpected Occurred."}
			utils.JSON(w, http.StatusInternalServerError, resp)
			log.Println(err.Error())
			return
		}

	}
}
