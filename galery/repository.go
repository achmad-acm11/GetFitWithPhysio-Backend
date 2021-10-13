package galery

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type RepositoryGalery interface {
	GetAll(ctx context.Context, tx *sql.Tx) []Galery
}

type repositoryGalery struct {
}

func NewRepositoryGalery() *repositoryGalery {
	return &repositoryGalery{}
}

func (r *repositoryGalery) GetAll(ctx context.Context, tx *sql.Tx) []Galery {
	query := "SELECT * FROM galeries"

	data, err := tx.QueryContext(ctx, query)
	helper.HandleError(err)
	defer data.Close()

	var galeries []Galery
	for data.Next() {
		galery := Galery{}

		data.Scan(&galery.Id, &galery.Photo, &galery.Caption, &galery.SubCaption)

		galeries = append(galeries, galery)
	}

	return galeries
}
