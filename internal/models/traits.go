package models

import (
	"log/slog"

	"github.com/oarielg/MasteriaBot/internal/database"
	"gorm.io/gorm"
)

type Trait struct {
	gorm.Model
	Name        string
	Description string `gorm:"type:TEXT"`
}

func ListTraits() ([]Trait, error) {
	var traits []Trait
	result := database.DB.Find(&traits)
	if result.Error != nil {
		slog.Error("error retreaving traits", "error", result.Error)
		return nil, result.Error
	}
	return traits, nil
}

func GetTrait(ID string) (Trait, error) {
	var trait Trait
	result := database.DB.First(&trait, ID)
	if result.Error != nil {
		slog.Error("error retreaving trait", "error", result.Error)
		return Trait{}, result.Error
	}
	return trait, nil
}

func AddCharacterTrait(char Character, trait Trait) error {
	result := database.DB.Model(&char).Association("Traits").Append(&trait)
	if result != nil {
		slog.Error("error adding trait to the character", "error", result)
		return result
	}
	return nil
}

func RemoveCharacterTrait(char Character, trait Trait) error {
	result := database.DB.Model(&char).Association("Traits").Delete(&trait)
	if result != nil {
		slog.Error("error removing trait from the character", "error", result)
		return result
	}
	return nil
}
