package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint
	PhoneNumber  string  `gorm:"type:varchar(20);uniqueIndex"`
	Email        *string `gorm:"type:varchar(30);uniqueIndex"`
	Token        string  `gorm:"size:500;uniqueIndex"`
	Name         string  `gorm:"type:varchar(20);"`
	Surname      string  `gorm:"type:varchar(25);"`
	ProfileImage string  `gorm:"size:500;"`
	CityID       *uint   `gorm:"type:bigint(20);index"`
	City         *City   `gorm:"foreignkey:city_id;references:id"`
	Documents    []Document
	JobStatus    string `gorm:"type:enum('trainee', 'agent', 'blocked');default:'trainee'"`
	Status       string `gorm:"type:enum('online', 'offline');default:'offline'"`
}
