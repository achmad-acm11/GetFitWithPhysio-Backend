package patient

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ServicePatient interface {
	GetAllService(ctx context.Context) []PatientResponse
}

type servicePatient struct {
	repo     ReposioryPatient
	db       *sql.DB
	validate *validator.Validate
}

func NewServicePatient(repo ReposioryPatient, db *sql.DB, validate *validator.Validate) *servicePatient {
	return &servicePatient{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

func (s *servicePatient) GetAllService(ctx context.Context) []PatientResponse {
	// Start Transaction
	tx, err := s.db.Begin()
	helper.HandleError(err)
	defer helper.CommitOrRollback(tx)

	patients := s.repo.GetAll(ctx, tx)

	return MapPatientsResponse(patients)
}
