package team

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryTeam interface {
	GetAll(cx context.Context, tx *gorm.DB) []Team
}

type repository struct {
}

// Function for access repo struct
func NewRepository() *repository {
	return &repository{}
}

// SQL Query Get All Data Team
func (r *repository) GetAll(cx context.Context, tx *gorm.DB) []Team {
	teams := []Team{}

	err := tx.WithContext(cx).Find(&teams).Error
	helper.HandleError(err)

	return teams
}
