package entity

import (
	"gorm.io/gorm"
)

type RoutePoint struct {
	gorm.Model
	ID      uint
	PointID uint
	RouteID uint
	Order   int
	Point   Point `gorm:"foreignkey:point_id;references:id"`
}
