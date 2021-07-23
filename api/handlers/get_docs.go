package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
Return "Hello World" when called.
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
