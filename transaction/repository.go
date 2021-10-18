package transaction

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	Create(ctx context.Context, tx *gorm.DB, transaction Transaction) Transaction
	Update(ctx context.Context, tx *gorm.DB, transaction Transaction) Transaction
}

type repositoryTransaction struct {
}

func NewRepositoryTransaction() *repositoryTransaction {
	return &repositoryTransaction{}
}

// Query SQL Create Transaction
func (r *repositoryTransaction) Create(ctx context.Context, tx *gorm.DB, transaction Transaction) Transaction {
	err := tx.WithContext(ctx).Create(&transaction).Error
	helper.HandleError(err)

	return transaction
}

// Query SQL Update Transaction
func (r *repositoryTransaction) Update(ctx context.Context, tx *gorm.DB, transaction Transaction) Transaction {
	err := tx.WithContext(ctx).Save(&transaction).Error
	helper.HandleError(err)

	return transaction
}
