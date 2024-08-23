package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name   string
	Traits []Trait `gorm:"many2many:character_traits;"`
	Powers []Power `gorm:"many2many:character_powers;"`
	Owner  string
}
