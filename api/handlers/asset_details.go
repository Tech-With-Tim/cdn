package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
)

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

func (s *Service) FetchAssetDetailsByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		id, err := strconv.Atoi(chi.URLParam(r, "path"))
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
