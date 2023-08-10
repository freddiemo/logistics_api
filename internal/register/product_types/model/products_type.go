package model

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type TransportationType uint8

const (
	TransportationTypeTerrestre TransportationType = 1
	TransportationTypeMaritime  TransportationType = 2
)

func (tt *TransportationType) Scan(value interface{}) error {
	b, ok := value.(uint8)
	if !ok {
		*tt = TransportationType(b)
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
	TransportationType TransportationType `json:"transportation_type" binding:"required" gorm:"type:transportation_type;default:1;column:transportation_type"`
	Code               string             `binding:"required" gorm:"not null"`
}
