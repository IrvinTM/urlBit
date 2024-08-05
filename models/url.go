package models

import (
	"fmt"
	"github.com/IrvinTM/urlBit/utils"
	"github.com/jinzhu/gorm"
)

type Url struct {
	gorm.Model
	Address   string    `json:"address"`
	ShortUrl  string    `json:"short_url"`
	UserId    uint      `json:"user_id"`
}

func (url *Url) Validate() (map[string] interface{}, bool) {

	if url.Address == "" {
		return utils.Message(false, "url  should be on the payload"), false
	}

	if url.UserId <= 0 {
		return utils.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return utils.Message(true, "success"), true
}

func (url *Url) Create() (map[string] interface{}) {

	if resp, ok := url.Validate(); !ok {
		return resp
	}

	GetDB().Create(url)

	resp := utils.Message(true, "success")
	resp["url"] = url
	return resp
}

func GetUrl(id uint) (*Url) {

	url := &Url{}
	err := GetDB().Table("urls").Where("id = ?", id).First(url).Error
	if err != nil {
		return nil
	}
	return url
}

func GetUrls(user uint) ([]*Url) {

	urls := make([]*Url, 0)
	err := GetDB().Table("urls").Where("user_id = ?", user).Find(&urls).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return urls
}

func GetByShortUrl(generatedUrl string) (*Url, error) {
    var url Url
    err := GetDB().Where("short_url = ?", generatedUrl).First(&url).Error
    if err != nil {
        return nil, err
    }
    return &url, nil
}