package promo

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryPromo interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Promo
	GetOneById(ctx context.Context, tx *gorm.DB, promoId int) Promo
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

func (r *repositoryPromo) GetOneById(ctx context.Context, tx *gorm.DB, promoId int) Promo {
	promo := Promo{}

	err := tx.WithContext(ctx).Where("id = ?", promoId).Find(&promo).Error
	helper.HandleError(err)

	return promo
}
