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
