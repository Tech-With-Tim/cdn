package api

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/dgrijalva/jwt-go"
)

type ctxString string

const errorstring string = "The server could not verify that you are authorized to access the URL requested. " +
	"You either supplied the wrong credentials (e.g. a bad password), " +
	"or your browser doesn't understand how to supply the credentials required."

func AuthJwtWrap(SecretKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var resp = map[string]interface{}{"error": "unauthorized", "message": "missing authorization token"}
			var header = r.Header.Get("Authorization")
			header = strings.TrimSpace(header)
			if header == "" {
				utils.JSON(w, http.StatusUnauthorized, resp)
				return
			}

			//utils.ExportVariables()
			token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
				return []byte(SecretKey), nil
			})

			if err != nil {
				resp["error"] = "unauthorized"
				if err.Error() == "Token is expired" {
					resp["message"] = err.Error()
					utils.JSON(w, http.StatusUnauthorized, resp)
					return
				}
				resp["message"] = errorstring
				utils.JSON(w, http.StatusUnauthorized, resp)
				log.Println(err.Error())
				return
			}

			claims, _ := token.Claims.(jwt.MapClaims)

			uid, err := strconv.Atoi(claims["uid"].(string))
			//fmt.Println(sub)
			if err != nil {
				resp["error"] = "something unexpected occurred"
				utils.JSON(w, http.StatusInternalServerError, resp)
				log.Println(err.Error())
				return
			}
			var ctxvalue ctxString = "uid"
			ctx := context.WithValue(r.Context(), ctxvalue, uid) // adding the user ID to the context
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
