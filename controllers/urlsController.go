package controllers

import (
	"net/http"

	"encoding/json"

	"strconv"

	"github.com/IrvinTM/urlBit/models"
	"github.com/IrvinTM/urlBit/utils"
	"github.com/gorilla/mux"
)

var CreateUrl = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	url := &models.Url{}

	err := json.NewDecoder(r.Body).Decode(url)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	url.UserId = user
	resp := url.Create()
	utils.Respond(w, resp)
}

var GetUrlsFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}
	
	data := models.GetUrl(uint(id))
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}