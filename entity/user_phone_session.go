package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserPhoneSession struct {
	gorm.Model
	ID          uint
	Token       string
	SMSCode     string
	PhoneNumber string
	IsUsed      int
	ExpireAt    time.Time
}
