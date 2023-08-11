package model

import (
	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Id      int    `json:"id"`
	Name    string `json:"name"`
	StateId int    `json:"state_id"`
}
