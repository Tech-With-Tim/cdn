package handlers

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
)

/*
Response: FileType | JSON

URL Parameters: AssetURL (string)

Return an asset, given the asset url. If the asset is not found,
the error message is returned along with the error status code. If it is
found, the asset is returned as per it's file type.
*/
func (s *Service) GetAsset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		var fileRow db.GetFileRow
		var err error

		url := chi.URLParam(r, "AssetUrl")

		fileRow, err = s.getFile(url, r.Context())

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
