package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/IrvinTM/urlBit/app"
	"github.com/IrvinTM/urlBit/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // adding the middleware

	router.HandleFunc("/api/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/newurl", controllers.CreateUrl).Methods("POST")
	router.HandleFunc("/api/myurls", controllers.GetUrlsFor).Methods("GET")
	router.HandleFunc("/{shorturl}", controllers.Redirect).Methods("GET")
	router.HandleFunc("/api/freeurl", controllers.CreateFreeUrl).Methods("POST")
	

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server running in port %s", port)

	err := http.ListenAndServe(":"+ port, router)

	if err != nil {
		fmt.Printf("\nThere was an error\n %v", err)
	}
}
