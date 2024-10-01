package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	database "github.com/kalleriakronos24/imbee-backend/db"
	"github.com/kalleriakronos24/imbee-backend/dto"
	"github.com/kalleriakronos24/imbee-backend/fb"
	"github.com/kalleriakronos24/imbee-backend/models"
)

type CheckExistingFcmJobStruct struct {
	*models.FcmJob
}

func (module *module) InsertFcmJob(p *dto.InsertFCM) (err error) {
	var response string
	if response, err = fb.SendFirebaseMessage(p.Message); err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	// if response from FCM not empty then we immediately save to database
	if response != "" {
		module.db.fcmJobModel.InsertFcmJob()
	}
	return err
}

func (module *module) RetrieveAllFcmJobPaginated(c *gin.Context, id uuid.UUID) (pagination *database.Pagination, err error) {
	return
}
