package entity

import "gorm.io/gorm"

type City struct {
	gorm.Model
	ID   uint
	Name string `gorm:"type:varchar(100)"`
}
