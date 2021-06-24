package handlers

import (
	"log"
	"net/http"
)

func HelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Header)
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			log.Println(err.Error())
		}
	}
}
