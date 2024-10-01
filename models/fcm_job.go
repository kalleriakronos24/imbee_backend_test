package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kalleriakronos24/imbee-backend/types"
	"github.com/kalleriakronos24/imbee-backend/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FcmJobOrm struct {
	db *gorm.DB
}

type FcmJob struct {
	ID         uuid.UUID `sql:"type:uuid;primary_key;" json:"id"`
	Identifier string    `json:"identifier" gorm:"not null" binding:"required"`
	DeliverAt  time.Time `json:"deliverAt"`
	types.DefaultModelProperty
}

func (base *FcmJob) BeforeCreate(scope *gorm.DB) error {
	uid, err := uuid.NewUUID()
	base.ID = uid
	base.Identifier = fmt.Sprintf("fmt-msg-%s", utils.RandStringBytes())
	return err
}

type FcmJobModelAction interface {
	GetOneByIdentifier(id string) (fcmJob FcmJob, err error)
	InsertFcmJob() (err error)
	DeleteFcmJob(id uuid.UUID, tx *gorm.DB) (err error)
}

func NewFcmJobAction(db *gorm.DB) FcmJobModelAction {
	return &FcmJobOrm{db}
}

func (o *FcmJobOrm) GetOneByIdentifier(id string) (fcmJob FcmJob, err error) {
	result := o.db.Model(&fcmJob).
		Where("identifier = ?", id).
		Preload(clause.Associations).
		First(&fcmJob)
	return fcmJob, result.Error
}

func (o *FcmJobOrm) InsertFcmJob() (err error) {
	result := o.db.Model(&FcmJob{}).Create(&FcmJob{})
	return result.Error
}

func (o *FcmJobOrm) DeleteFcmJob(id uuid.UUID, tx *gorm.DB) (err error) {
	result := tx.Model(&FcmJob{}).Delete(&FcmJob{}, id)
	return result.Error
}
