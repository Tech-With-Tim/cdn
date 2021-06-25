package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON returns a well formated response with a status code
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		_, _ = fmt.Fprintf(w, "%s", err.Error())
	}
}
