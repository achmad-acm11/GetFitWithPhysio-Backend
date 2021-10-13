package promo

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type RepositoryPromo interface {
	GetAll(ctx context.Context, tx *sql.Tx) []Promo
}

type repositoryPromo struct {
}

func NewRepositoryPromo() *repositoryPromo {
	return &repositoryPromo{}
}

func (r *repositoryPromo) GetAll(ctx context.Context, tx *sql.Tx) []Promo {
	query := "SELECT * FROM promos"

	data, err := tx.QueryContext(ctx, query)

	helper.HandleError(err)

	var promos []Promo
	for data.Next() {
		promo := Promo{}
		err := data.Scan(&promo.Id, &promo.Discount)
		helper.HandleError(err)

		promos = append(promos, promo)
	}
	return promos
}
