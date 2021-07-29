package handlers

import (
	"database/sql"
	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

/*
Response: JSON

URL Parameters: path (String)

Returns details of assets, given the path to the asset.
*/
func FetchAssetDetailsByURL(store *db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		url := chi.URLParam(r, "path")
		fileRow, err := store.GetAssetDetailsByUrl(r.Context(), url)
		// type GetAssetDetailsByUrlRow struct {
		//	ID   int64  `json:"id"`
		//	Name string `json:"name"`
		//  CreatorID int64  `json:"creatorID"`
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

/*
Response: JSON

URL Parameters: id (Integer)

Returns details of assets, given the asset id. If it finds the asset,
it returns JSON containing info about the asset. If it doesn't, it
returns the error along with the error code.
*/
func FetchAssetDetailsByID(store *db.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			resp = map[string]interface{}{"error": "id is not a valid integer."}
			utils.JSON(w, http.StatusBadRequest, resp)
			log.Println(err.Error())
			return
		}

		fileRow, err := store.GetAssetDetailsById(r.Context(), int64(id))
		// type GetAssetDetailsByIdRow struct {
		//    UrlPath   string `json:"urlPath"`
		//    Name      string `json:"name"`
		//    CreatorID int64  `json:"creatorID"`
		//	}
		if err != nil {
			if err == sql.ErrNoRows {
				resp = map[string]interface{}{"error": "not found",
					"message": "no asset found with that id."}
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
