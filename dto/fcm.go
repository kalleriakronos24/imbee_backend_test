package dto

type InsertFCM struct {
	Message string `json:"message" validate:"required"`
}
