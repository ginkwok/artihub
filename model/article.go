package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `gorm:"type:VARCHAR(255);"`
	Content string `gorm:"type:TEXT;"`
}
