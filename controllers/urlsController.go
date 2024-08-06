package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/IrvinTM/urlBit/models"
	"github.com/IrvinTM/urlBit/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var CreateUrl = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	url := &models.Url{}

	err := json.NewDecoder(r.Body).Decode(url)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}

	var shortUrl string
	for{
		shortUrl = utils.GenShort()
		//TODO check if the regular url already exists
		if url, err := models.GetByShortUrl(shortUrl); url != nil {
			if err != nil {
				utils.Respond(w, utils.Message(false, "Error while checking the url"))
				return
			}
		}else{
			break
		}
	}
	if url1, err := models.GetByRegularUrl(url.Address); url1 != nil {
		if err != nil {

		}else{
		utils.Respond(w, utils.Message(false, "Url already registered"))
		return
		}
	}


	url.ShortUrl = shortUrl
	url.UserId = user
	resp := url.Create()
	utils.Respond(w, resp)
}

var GetUrlsFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetUrls(id)
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}

var Redirect = func(w http.ResponseWriter, r *http.Request) {
    // Extract the shorturl from the request URL
    vars := mux.Vars(r)
    shortURL := vars["shorturl"]

    // Query the database to find the original URL
    var url models.Url
    err := models.GetDB().Where("short_url = ?", shortURL).First(&url).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "URL not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Redirect the user to the original URL
    http.Redirect(w, r, url.Address, http.StatusMovedPermanently)
}