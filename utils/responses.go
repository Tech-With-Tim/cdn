package utils

import (
	"encoding/json"
	"net/http"
)

// JSON returns a well formated response with a status code
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(500)
		er := json.NewEncoder(w).Encode(map[string]interface{}{"error": "something unexpected occurred."})
		if er != nil {
			return
		}
	}
}
