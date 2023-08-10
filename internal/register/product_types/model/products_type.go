package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type TransportationType uint8

const (
	TransportationTypeTerrestre TransportationType = 1
	TransportationTypeMaritime  TransportationType = 2
)

func (tt *TransportationType) Scan(value interface{}) error {
	switch value.(type) {
	case string:
		v2, err := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 8)
		if err != nil {
			fmt.Println(err.Error())
		}
		*tt = TransportationType(v2)
	}

	return nil
}

func (tt TransportationType) Value() (driver.Value, error) {
	return uint8(tt), nil
}

func (tt TransportationType) IsValid() bool {
	switch tt {
	case 1, 2:
		return true
	}

	return false
}

type ProductType struct {
	gorm.Model
	Name               string             `binding:"required" gorm:"type:varchar(32);not null;UNIQUE"`
	TransportationType TransportationType `json:"transportation_type" binding:"required" gorm:"type:transportation_type;column:transportation_type"`
	Code               string             `binding:"required" gorm:"not null"`
}
