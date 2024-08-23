package migrations

import (
	"github.com/oarielg/MasteriaBot/internal/database"
	"github.com/oarielg/MasteriaBot/internal/models"
)

func AutoMigrate() {
	database.DB.AutoMigrate(models.Trait{}, models.Power{}, models.Character{})
	MigrateTraits()
	MigratePowers()
}
