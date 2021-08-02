package entity

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProfileLog struct {
	gorm.Model
	ID        uint
	ProfileID uint
	Action    string `gorm:"type:enum('create', 'update', 'delete')"`
	Changer   string `gorm:"type:enum('user', 'admin')"`
	Data      datatypes.JSON
}
