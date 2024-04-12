package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	*gorm.Model
	ID    uint `gorm:"primaryKey"`
	Title string
	Text  string
}
