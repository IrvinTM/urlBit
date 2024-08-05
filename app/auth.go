package app

import (
	"context"
	"fmt"
	// "fmt"
	"net/http"
	"os"
	"strings"

	"github.com/IrvinTM/urlBit/models"
	"github.com/IrvinTM/urlBit/utils"
	"github.com/dgrijalva/jwt-go"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuth := []string{"/register", "/login"} //noauth endpoints
		requestPath := r.URL.Path                              // current request path

		// check if the req needs auth and serve it if it doesnt
		for _, value := range noAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		// if we require auth we continue the execution

		response := make(map[string]interface{})     // i may not need this one
		tokenHeader := r.Header.Get("Authorization") // get the token from the request header

		//check if the token is empty and if it is "" return 403
		if tokenHeader == "" {
			response = utils.Message(false, "Invalid or malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid or malformed auth token")
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		tokenPart := splitted[1] // get the token part in the second index wich is the one we need
		tk := &models.Token{}
		fmt.Printf("this is the token %s",splitted[1])

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utils.Message(false, "Malformed authentication token 1")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		
		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}
		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		// fmt.Sprintf("User %v", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
