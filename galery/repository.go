package galery

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryGalery interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Galery
	Create(ctx context.Context, tx *gorm.DB, galery Galery) Galery
	UploadPhoto(ctx context.Context, tx *gorm.DB, galeryId int, filePath string)
}

type repositoryGalery struct {
}

func NewRepositoryGalery() *repositoryGalery {
	return &repositoryGalery{}
}

// Query SQL Get All Data
func (r *repositoryGalery) GetAll(ctx context.Context, tx *gorm.DB) []Galery {
	galeries := []Galery{}
	err := tx.WithContext(ctx).Find(&galeries).Error
	helper.HandleError(err)

	return galeries
}

// Query SQL Create Galery
func (r *repositoryGalery) Create(ctx context.Context, tx *gorm.DB, galery Galery) Galery {
	err := tx.WithContext(ctx).Create(&galery).Error
	helper.HandleError(err)

	return galery
}

// Query SQL Upload Foto
func (r *repositoryGalery) UploadPhoto(ctx context.Context, tx *gorm.DB, galeryId int, filePath string) {
	err := tx.WithContext(ctx).Model(Galery{}).Where("id = ?", galeryId).Update("photo", filePath).Error
	helper.HandleError(err)
}
