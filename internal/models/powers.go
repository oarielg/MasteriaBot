package models

import (
	"log/slog"

	"github.com/oarielg/MasteriaBot/internal/database"
	"gorm.io/gorm"
)

type Power struct {
	gorm.Model
	Name        string
	Description string `gorm:"type:TEXT"`
	Willpower   string
	Duration    string
}

func ListPowers() ([]Power, error) {
	var powers []Power
	result := database.DB.Find(&powers)
	if result.Error != nil {
		slog.Error("error retreaving powers", "error", result.Error)
		return nil, result.Error
	}
	return powers, nil
}

func GetPower(ID string) (Power, error) {
	var power Power
	result := database.DB.First(&power, ID)
	if result.Error != nil {
		slog.Error("error retreaving power", "error", result.Error)
		return Power{}, result.Error
	}
	return power, nil
}

func AddCharacterPower(char Character, power Power) error {
	result := database.DB.Model(&char).Association("Powers").Append(&power)
	if result != nil {
		slog.Error("error adding power to the character", "error", result)
		return result
	}
	return nil
}

func RemoveCharacterPower(char Character, power Power) error {
	result := database.DB.Model(&char).Association("Powers").Delete(&power)
	if result != nil {
		slog.Error("error removing power from the character", "error", result)
		return result
	}
	return nil
}
