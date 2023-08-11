package model

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CountryId int    `json:"country_id"`
}
