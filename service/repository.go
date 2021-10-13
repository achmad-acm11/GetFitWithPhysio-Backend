package service

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type RepositoryService interface {
	GetAll(ctx context.Context, tx *sql.Tx) []Service
}

type repositoryService struct {
}

func NewRepositoryService() *repositoryService {
	return &repositoryService{}
}

func (r *repositoryService) GetAll(ctx context.Context, tx *sql.Tx) []Service {
	query := "SELECT * FROM services"

	data, err := tx.QueryContext(ctx, query)
	helper.HandleError(err)
	defer data.Close()

	var services []Service
	for data.Next() {
		service := Service{}

		err := data.Scan(&service.Id, &service.Kode_promo, &service.Service_name, &service.Image, &service.Description)
		helper.HandleError(err)

		services = append(services, service)
	}

	return services
}
