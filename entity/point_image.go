package entity

import "gorm.io/gorm"

type PointImage struct {
	gorm.Model
	ID      uint
	PointID uint
	Path    string
}
