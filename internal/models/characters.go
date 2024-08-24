package models

import (
	"log/slog"

	"github.com/oarielg/MasteriaBot/internal/database"
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name      string
	Vitality  int     `gorm:"default:15"`
	Willpower int     `gorm:"default:15"`
	Traits    []Trait `gorm:"many2many:character_traits;"`
	Powers    []Power `gorm:"many2many:character_powers;"`
	Owner     string
}

func CreateCharacter(name, owner string) error {
	char := Character{Name: name, Owner: owner}
	result := database.DB.Create(&char)
	if result.Error != nil {
		slog.Error("error creating character", "error", result.Error)
		return result.Error
	}
	return nil
}

func ListCharacters(owner string) ([]Character, error) {
	var chars []Character
	result := database.DB.Where("owner = ?", owner).Find(&chars)
	if result.Error != nil {
		slog.Error("error retreaving characters", "error", result.Error)
		return nil, result.Error
	}
	return chars, nil
}

func GetCharacter(ID string) (Character, error) {
	var char Character
	result := database.DB.Preload("Traits").Preload("Powers").First(&char, ID)
	if result.Error != nil {
		slog.Error("error retreaving characters", "error", result.Error)
		return Character{}, result.Error
	}
	return char, nil
}
