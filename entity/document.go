package entity

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	ID     uint
	UserID uint
	Path   string
}
