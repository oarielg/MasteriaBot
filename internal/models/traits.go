package models

import "gorm.io/gorm"

type Trait struct {
	gorm.Model
	Name        string
	Description string
}
