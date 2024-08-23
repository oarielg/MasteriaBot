package models

import "gorm.io/gorm"

type Power struct {
	gorm.Model
	Name        string
	Description string
	Willpower   string
	Duration    string
}
