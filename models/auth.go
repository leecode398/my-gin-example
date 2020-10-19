package models

import "github.com/jinzhu/gorm"

type Auth struct {
	ID       int    `gorm:"primary_key" jsonL"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return true, nil
}
