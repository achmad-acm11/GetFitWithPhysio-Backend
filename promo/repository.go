package promo

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryPromo interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Promo
}

type repositoryPromo struct {
}

func NewRepositoryPromo() *repositoryPromo {
	return &repositoryPromo{}
}

func (r *repositoryPromo) GetAll(ctx context.Context, tx *gorm.DB) []Promo {
	promos := []Promo{}
	err := tx.WithContext(ctx).Find(&promos).Error
	helper.HandleError(err)

	return promos
}
