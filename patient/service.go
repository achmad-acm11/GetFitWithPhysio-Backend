package patient

import (
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/user"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ServicePatient interface {
	GetAllService(ctx context.Context) []PatientResponse
	Register(ctx context.Context, req RegisterRequest) RegisterResponse
}

type servicePatient struct {
	repo     ReposioryPatient
	repoUser user.RepositoryUser
	db       *sql.DB
	validate *validator.Validate
}

func NewServicePatient(repo ReposioryPatient, repoUser user.RepositoryUser, db *sql.DB, validate *validator.Validate) *servicePatient {
	return &servicePatient{
		repo:     repo,
		repoUser: repoUser,
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

func (s *servicePatient) Register(ctx context.Context, req RegisterRequest) RegisterResponse {
	// Validate Format Request
	err := s.validate.Struct(req)
	helper.HandleError(err)

	// Start Transaction
	tx, err := s.db.Begin()
	helper.HandleError(err)
	defer helper.CommitOrRollback(tx)

	// Create Data User
	dataUser := s.repoUser.Create(ctx, tx, user.User{
		Role:     1,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	// Create Data Patient
	dataPatient := s.repo.Create(ctx, tx, Patient{
		Id_user: dataUser.Id,
		Gender:  req.Gender,
		Nik:     req.Nik,
		Phone:   req.Phone,
		Address: req.Address,
	})

	return MapRegisterResponse(dataPatient, dataUser)
}
