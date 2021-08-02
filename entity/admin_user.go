package entity

import "golang.org/x/crypto/bcrypt"

type AdminUser struct {
	ID       uint
	Username string `gorm:"type:varchar(255);"`
	Password string `gorm:"type:varchar(255);"`
}

func (adminUser *AdminUser) SetPasswordHash() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(adminUser.Password), bcrypt.DefaultCost)
	adminUser.Password = string(bytes)
}
