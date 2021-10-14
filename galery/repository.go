package galery

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryGalery interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Galery
}

type repositoryGalery struct {
}

func NewRepositoryGalery() *repositoryGalery {
	return &repositoryGalery{}
}

func (r *repositoryGalery) GetAll(ctx context.Context, tx *gorm.DB) []Galery {
	galeries := []Galery{}
	err := tx.WithContext(ctx).Find(&galeries).Error
	helper.HandleError(err)

	return galeries
}
