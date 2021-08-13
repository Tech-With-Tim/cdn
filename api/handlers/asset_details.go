package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
)

/*
Response: JSON

URL Parameters: path (String)

Description: "Returns details of assets, given the path to the asset. If
the asset is not found, a 404 error is raised"
*/
func (s *Service) FetchAssetDetailsByURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		url := chi.URLParam(r, "path")
		fileRow, err := s.Store.GetAssetDetailsByUrl(r.Context(), url)

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

Description: "Returns details of assets, given the asset ID.
If it finds the asset, it returns JSON containing info about
the asset. If the asset is not found, a 404 error is raised.
If the ID provided is not an integer, a 400 error is raised."
*/
func (s *Service) FetchAssetDetailsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			resp = map[string]interface{}{"error": "id is not a valid integer."}
			utils.JSON(w, http.StatusBadRequest, resp)
			log.Println(err.Error())
			return
		}

		fileRow, err := s.Store.GetAssetDetailsById(r.Context(), int64(id))
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
