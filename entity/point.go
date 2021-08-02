package entity

import "gorm.io/gorm"

type Point struct {
	gorm.Model
	ID           uint
	UserID       uint    `gorm:"type:bigint(20)"`
	Longitude    float64 `gorm:"type:decimal(14, 9)"`
	Latitude     float64 `gorm:"type:decimal(14, 9)"`
	PointImages  []PointImage
	CityID       *uint  `gorm:"type:bigint(20);index"`
	City         *City  `gorm:"foreignkey:city_id;references:id"`
	Comment      string `gorm:"type:text"`
	Type         string `gorm:"type:enum('shop', 'supermarket', 'gas_station');default:'shop'"`
	PhoneNumber  string `gorm:"type:varchar(100)"`
	Address      string `gorm:"type:text"`
	BusinessName string `gorm:"type:text"`
}
