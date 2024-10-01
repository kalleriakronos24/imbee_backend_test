package migrations

import (
	"fmt"

	database "github.com/kalleriakronos24/imbee-backend/db"
	"github.com/kalleriakronos24/imbee-backend/models"
)

func Migrate() {
	// auto migration all models
	if err := database.GetDatabaseConnection().AutoMigrate(
		models.FcmJob{},
	); err != nil {
		panic(err)
	}
	fmt.Println("Database Migrated")
}
