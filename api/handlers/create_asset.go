package handlers

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strconv"

	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/omeid/pgerror"
	"golang.org/x/sync/errgroup"
)

/*
getUrlPath generates a random url path if url path is not provided
in the request.
*/
func getUrlPath(url string, fileExt []string) string {
	if url == "" {
		if len(fileExt) != 0 {
			if fileExt[0] != "" {
				return utils.RandomString(16) + fileExt[0]
			}
			return utils.RandomString(16)
		}
		return utils.RandomString(16)
	}
	return url

}

// storeAsset Creates and stores asset in the database
func storeAsset(mimetype string,
	fileName string,
	fileData []byte,
	assetName string,
	assetUrlPath string,
	userId int,
	store *db.Store,
	w http.ResponseWriter,
	r *http.Request) (assetId int64, err error) {
	var resp map[string]interface{}
	params := db.CreateAssetParams{
		Mimetype:  mimetype,
		Name:      fileName,
		Data:      fileData,
		Name_2:    assetName, //Asset Name
		UrlPath:   assetUrlPath,
		CreatorID: int64(userId),
	}
	assetId, err = store.CreateAssetFile(r.Context(), params)
	if err != nil {
		if e := pgerror.UniqueViolation(err); e != nil {
			resp = map[string]interface{}{"error": "Conflict",
				"message": "Asset with this url_path already exists."}
			utils.JSON(w, http.StatusConflict, resp)
			return
		}
		resp = map[string]interface{}{"error": "Something unexpected occurred."}
		utils.JSON(w, http.StatusInternalServerError, resp)
		log.Println(err.Error())
		return
	}
	return
}

/*
Response: JSON

URL Parameters: None

Request Body : [
	{name: "name", type: "string", description: "the name of the asset"},
]
// we can basically create a list like this which would be in the docs.json file
then we can render this list in a table with svelte
Create Asset creates an asset with a given file, uploaded
under the `data` parameter in the form data. If it succeeds,
it returns a 201 Created Status. If the file is too large, a
413 error is raised. If the file is not provided, a 400
error is raised.
*/
func (s *Service) CreateAsset(FileSize int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp map[string]interface{}
		fileData := make(chan []byte)
		var assetId int64
		var urlPath string
		g, ctx := errgroup.WithContext(r.Context())

		r.Body = http.MaxBytesReader(w, r.Body, FileSize<<20)
		//Parse Form
		err := r.ParseMultipartForm(FileSize << 20)
		if err != nil {
			resp = map[string]interface{}{"error": err.Error()}
			utils.JSON(w, http.StatusRequestEntityTooLarge, resp)
			return
		}

		upload, handler, err := r.FormFile("data")
		if err != nil {
			resp = map[string]interface{}{"error": "No file in 'data' field"}
			utils.JSON(w, http.StatusBadRequest, resp)
			return
		}
		defer func(upload multipart.File) {
			err := upload.Close()
			if err != nil {
				log.Println(err.Error())
			}
		}(upload)

		g.Go(func() error {
			defer close(fileData)
			fileBytes, er := ioutil.ReadAll(upload)
			if er != nil {
				log.Println(er.Error())
				resp = map[string]interface{}{"error": "Something unexpected occurred."}
				utils.JSON(w, http.StatusInternalServerError, resp)
				return er
			}
			fileData <- fileBytes
			return nil

		})
		//var params db.CreateAssetParams
		g.Go(func() error {
			var fileExt []string

			fileName := handler.Filename

			assetName := r.FormValue("name")
			mimetype := handler.Header.Get("Content-Type")
			//File Extension
			fileExt, _ = mime.ExtensionsByType(mimetype)
			userId := ctx.Value("uid").(int)
			urlPath = getUrlPath(html.EscapeString(r.FormValue("url_path")), fileExt)
			//Storing asset in database
			assetId, err = storeAsset(mimetype, fileName, <-fileData,
				assetName, urlPath, userId, s.Store, w, r)
			if err != nil {
				return err
			}
			return nil
		})

		if err = g.Wait(); err != nil {
			return
		}
		resp = map[string]interface{}{"location": fmt.Sprintf("%s/%s", r.Host, urlPath),
			"asset_id": strconv.Itoa(int(assetId))}

		utils.JSON(w, http.StatusCreated, resp)

	}
}
