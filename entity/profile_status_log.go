package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProfileStatusLog struct {
	gorm.Model
	ID        uint
	UserID    uint
	OnlineAt  time.Time `gorm:"type:time"`
	OfflineAt time.Time `gorm:"type:time"`
	Track     string    `gorm:"type:text"`
}
