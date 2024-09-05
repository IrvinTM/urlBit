package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/IrvinTM/urlBit/models"
	"github.com/IrvinTM/urlBit/utils"
	"github.com/dgrijalva/jwt-go"
)

var CORSMiddleware = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        allowedOrigins := []string{
            "https://shortener.archbtw.site",
            "http://localhost:2323",
            "http://localhost:80",
            "http://localhost:2323",
            "https://archbtw.site",
            "http://localhost:3000",

        }

        origin := r.Header.Get("Origin")
        fmt.Printf("this is the origin :%v",origin)
        for _, allowedOrigin := range allowedOrigins {
            if origin == allowedOrigin {
                w.Header().Set("Access-Control-Allow-Origin", origin)
                break
            }
        }

        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        // If it's a preflight request, return 200
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

type contextKey string

func (c contextKey) String() string {
    return "app context key " + string(c)
}

const userKey = contextKey("user")

var JwtAuthentication = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        noAuth := []string{"/api/register", "/api/login", "/{shorturl}", "/api/freeurl"} // no-auth endpoints
        requestPath := r.URL.Path                              // current request path

        // check if the req needs auth and serve it if it doesn't
        for _, value := range noAuth {
            if value == requestPath {
                next.ServeHTTP(w, r)
                return
            }
        }
        // if it is a short url
        if strings.HasPrefix(requestPath, "/") && len(strings.Split(requestPath, "/")) == 2 {
            next.ServeHTTP(w, r)
            return
        }

        // Handle JWT auth
        tokenHeader := r.Header.Get("Authorization") // get the token from the request header

        if tokenHeader == "" {
            response := utils.Message(false, "Missing or malformed auth token")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            utils.Respond(w, response)
            return
        }

        splitted := strings.Split(tokenHeader, " ")
        if len(splitted) != 2 {
            response := utils.Message(false, "Invalid or malformed auth token")
            w.Header().Add("Content-Type", "application/json")
            utils.Respond(w, response)
            return
        }

        tokenPart := splitted[1]
        tk := &models.Token{}

        token, err := jwt.ParseWithClaims(tokenPart, tk, func(t *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("token_password")), nil
        })

        if err != nil || !token.Valid {
            response := utils.Message(false, "Malformed or expired authentication token")
            w.WriteHeader(http.StatusForbidden)
            w.Header().Add("Content-Type", "application/json")
            utils.Respond(w, response)
            return
        }

        // Set the user in context
        ctx := context.WithValue(r.Context(), userKey, tk.UserId)
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
    })
}

// Chain the middlewares
func ChainMiddlewares(handler http.Handler) http.Handler {
    return CORSMiddleware(JwtAuthentication(handler))
}
