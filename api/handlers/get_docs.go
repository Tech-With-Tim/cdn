package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tech-With-Tim/cdn/utils"
)

/*
Response: JSON

URL Parameters: None

Returns documentation for the CDN routes. Each route has
information regarding the response type, methods allowed,
URL parameters, possible errors, and a brief description
of the route.
*/
func GetDocs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		file, err := os.Open("./docs/docs.json")

		if err != nil {
			resp := map[string]interface{}{"error": "Something Unexpected Occurred."}
			utils.JSON(w, http.StatusInternalServerError, resp)
		}

		defer file.Close()
		data, _ := ioutil.ReadAll(file)

		_, err = w.Write(data)
		if err != nil {
			resp := map[string]interface{}{"error": "Something Unexpected Occurred."}
			utils.JSON(w, http.StatusInternalServerError, resp)
		}
	}
}

// Serve the docs page
func ServeDocsPage() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
	    
		path := r.URL.Path
		path = filepath.Join("./docs/docs-template/public", path)
		_, err := os.Stat(path)

		if os.IsNotExist(err) {

			http.ServeFile(w, r, "./docs/docs-template/public/index.html")
			return

		} else if err != nil {
			
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		
		}

		http.FileServer(http.Dir("./docs/docs-template/public")).ServeHTTP(w, r)
	}
}