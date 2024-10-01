package services

import (
	database "github.com/kalleriakronos24/imbee-backend/db"
	"github.com/kalleriakronos24/imbee-backend/dto"
	"github.com/kalleriakronos24/imbee-backend/models"
	"gorm.io/gorm"
)

var Handler HandlerFunc

type HandlerFunc interface {
	// FCM
	InsertFcmJob(p *dto.InsertFCM) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn        *gorm.DB
	fcmJobModel models.FcmJobModelAction
}

type GenerateDocumentOutput struct {
	OutputPath string
	FileName   string
}

func InitializeServices() (err error) {
	// Initialize DB
	db := database.GetDatabaseConnection()

	Handler = &module{
		db: &dbEntity{
			conn:        db,
			fcmJobModel: models.NewFcmJobAction(db),
		},
	}
	return
}
