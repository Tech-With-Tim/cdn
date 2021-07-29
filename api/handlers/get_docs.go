package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

/*
Response: JSON

URL Parameters: None

Returns documentation for the CDN routes.
*/
func GetDocs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		file, err := os.Open("./docs/docs.json")

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error while reading docs"))
		}

		defer file.Close()
		data, _ := ioutil.ReadAll(file)

		_, err = w.Write(data)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

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