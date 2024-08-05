package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IrvinTM/urlBit/app"
	"github.com/IrvinTM/urlBit/controllers"
	"github.com/IrvinTM/urlBit/utils"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // adding the middleware

	router.HandleFunc("/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/newurl", controllers.CreateUrl).Methods("POST")
	router.HandleFunc("/myurls", controllers.GetUrlsFor).Methods("GET")
	port := os.Getenv("PORT")

	fmt.Printf("the random url is %s \n", utils.GenShort())

	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server running in port %s", port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Printf("There was an error %v", err)
	}

}
