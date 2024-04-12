package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID       uint `gorm:"primaryKey"`
	Name     string
	Username string `gorm:"uniqueIndex"`
	Password string
	Admin    bool `gorm:"default:false"`
}
