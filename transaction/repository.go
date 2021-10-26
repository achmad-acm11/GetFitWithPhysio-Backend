package transaction

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"gorm.io/gorm"
)

type RepositoryTransaction interface {
	GetAll(ctx context.Context, tx *gorm.DB) []Transaction
	Create(ctx context.Context, tx *gorm.DB, transaction Transaction) Transaction
	Update(ctx context.Context, tx *gorm.DB, transaction Transaction) Transaction
}

type repositoryTransaction struct {
}

func NewRepositoryTransaction() *repositoryTransaction {
	return &repositoryTransaction{}
}

// Query SQL Get All Transaction
func (r *repositoryTransaction) GetAll(ctx context.Context, tx *gorm.DB) []Transaction {
	result := []map[string]interface{}{}
	err := tx.Table("transactions").Joins("join users ON users.id = transactions.id_user").Joins("join services ON services.id = transactions.id_service").Joins("join patients ON patients.id_user = users.id").Find(&result).Error
	helper.HandleError(err)

	transactions := MapTransactions(result)

	return transactions
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
