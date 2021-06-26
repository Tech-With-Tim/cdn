package handlers

import (
	"database/sql"
	db "github.com/Ibezio/cdn/db/sqlc"
	"github.com/Ibezio/cdn/utils"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func FetchAssetDetails(store *db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		url := chi.URLParam(r, "path")
		fileRow, err := store.GetAssetDetailsByUrl(r.Context(), url)
		// type GetAssetDetailsByUrlRow struct {
		//	ID   int64  `json:"id"`
		//	Name string `json:"name"`
		// }
		if err != nil {
			if err == sql.ErrNoRows {
				resp = map[string]interface{}{"error": "not found",
					"message": "no asset found with that url path."}
				utils.JSON(w, http.StatusNotFound, resp)
				return
			}
			resp = map[string]interface{}{"error": "something Unexpected Occurred."}
			utils.JSON(w, http.StatusInternalServerError, resp)
			log.Println(err.Error())
			return
		}
		utils.JSON(w, http.StatusOK, fileRow)
	}
}