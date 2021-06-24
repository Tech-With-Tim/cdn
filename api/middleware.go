package api

import (
	"context"
	"github.com/Ibezio/cdn/utils"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const errorstring string = "The server could not verify that you are authorized to access the URL requested. " +
	"You either supplied the wrong credentials (e.g. a bad password), " +
	"or your browser doesn't understand how to supply the credentials required."

type contextSub string

func AuthJwtWrap(SecretKey, issuer string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var substr contextSub = "sub"
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

			if claims["iss"].(string) != issuer { // "Ibezio Development" {
				resp["error"] = "unauthorized"
				resp["message"] = "jwt token issuer is invalid. try regenerating the token"
				utils.JSON(w, http.StatusUnauthorized, resp)
				return
			}
			sub, err := strconv.Atoi(claims["sub"].(string))
			//fmt.Println(sub)
			if err != nil {
				resp["error"] = "something unexpected occurred"
				resp["message"] = err.Error()
				utils.JSON(w, http.StatusInternalServerError, resp)
				log.Println(err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), substr, sub) // adding the user ID to the context
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
