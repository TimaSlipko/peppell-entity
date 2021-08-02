package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type District struct {
	gorm.Model
	ID          uint
	Name        string `gorm:"type:text"`
	Okato       string `gorm:"type:text"`
	Abbrev      string `gorm:"type:text"`
	Type        string `gorm:"type:text"`
	Coordinates datatypes.JSON
}
