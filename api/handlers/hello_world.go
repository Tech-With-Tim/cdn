package handlers

import (
	"log"
	"net/http"
)

/*
Response: String

URL Parameters: None

Returns `Hello, World!` when called. This route is for
testing purposes only.
*/
func HelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Header)
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			log.Println(err.Error())
		}
	}
}