package team

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryTeam interface {
	GetAll(cx context.Context, tx *gorm.DB) []Team
	Create(ctx context.Context, tx *gorm.DB, team Team) Team
	UpdatePhoto(ctx context.Context, tx *gorm.DB, teamId int, filePath string)
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

// SQL Query Create Team
func (r *repository) Create(ctx context.Context, tx *gorm.DB, team Team) Team {

	err := tx.WithContext(ctx).Create(&team).Error
	helper.HandleError(err)

	return team
}

// SQL Query Update Photo
func (r *repository) UpdatePhoto(ctx context.Context, tx *gorm.DB, teamId int, filePath string) {
	err := tx.WithContext(ctx).Model(Team{}).Where("id = ?", teamId).Update("photo_team", filePath).Error
	helper.HandleError(err)
}
