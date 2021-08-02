package entity

import (
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	ID           uint
	Name         string `gorm:"type:text"`
	UserID       *uint  `gorm:"type:bigint(20);index:idx_unique"`
	User         User   `gorm:"foreignkey:user_id;references:id"`
	Route        string `gorm:"type:text"`
	Description  string `gorm:"type:text"`
	District     string `gorm:"type:text"`
	Length       int
	PointsAmount int
	Day          int
	CityID       *uint `gorm:"type:bigint(20);index:idx_unique"`
	City         City  `gorm:"foreignkey:city_id;references:id"`
}
