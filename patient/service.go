package patient

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/user"
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServicePatient interface {
	GetAllService(ctx context.Context) []PatientResponse
	Register(ctx context.Context, req RegisterRequest) PatientResponse
	CreateService(ctx context.Context, req CreatePatientRequest) PatientResponse
	DetailService(ctx context.Context, patientId int) PatientResponse
}

type servicePatient struct {
	repo     ReposioryPatient
	repoUser user.RepositoryUser
	db       *gorm.DB
	validate *validator.Validate
}

func NewServicePatient(repo ReposioryPatient, repoUser user.RepositoryUser, db *gorm.DB, validate *validator.Validate) *servicePatient {
	return &servicePatient{
		repo:     repo,
		repoUser: repoUser,
		db:       db,
		validate: validate,
	}
}

func (s *servicePatient) GetAllService(ctx context.Context) []PatientResponse {
	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	patients := s.repo.GetAll(ctx, tx)

	return MapPatientsResponse(patients)
}

func (s *servicePatient) Register(ctx context.Context, req RegisterRequest) PatientResponse {
	// Validate Format Request
	err := s.validate.Struct(req)
	helper.HandleError(err)

	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	// Create Data User
	dataUser := s.repoUser.Create(ctx, tx, user.User{
		Role:     1,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	date, _ := time.Parse("02-01-2006", req.Birth_date)

	// Create Data Patient
	dataPatient := s.repo.Create(ctx, tx, Patient{
		Id_user:    dataUser.Id,
		Gender:     req.Gender,
		Nik:        req.Nik,
		Birth_date: date,
		Phone:      req.Phone,
		Address:    req.Address,
	})

	dataPatient.User.Email = dataUser.Email
	dataPatient.User.Role = dataUser.Role
	dataPatient.User.Name = dataUser.Name
	dataPatient.User.Photo_user = dataUser.Photo_user

	return MapPatientResponse(dataPatient)
}

func (s *servicePatient) CreateService(ctx context.Context, req CreatePatientRequest) PatientResponse {
	// Validate Format Request
	err := s.validate.Struct(req)
	helper.HandleError(err)

	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	// Create Data User
	dataUser := s.repoUser.Create(ctx, tx, user.User{
		Role:     1,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	date, _ := time.Parse("02-01-2006", req.Birth_date)

	// Create Data Patient
	dataPatient := s.repo.Create(ctx, tx, Patient{
		Id_user:    dataUser.Id,
		Gender:     req.Gender,
		Nik:        req.Nik,
		Birth_date: date,
		Phone:      req.Phone,
		Address:    req.Address,
	})

	dataPatient.User.Email = dataUser.Email
	dataPatient.User.Role = dataUser.Role
	dataPatient.User.Name = dataUser.Name
	dataPatient.User.Photo_user = dataUser.Photo_user

	return MapPatientResponse(dataPatient)
}

func (s *servicePatient) DetailService(ctx context.Context, patientId int) PatientResponse {
	// Star Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	patient := s.repo.GetOneById(ctx, tx, patientId)

	if patient.Id == 0 {
		panic(exception.NewNotFoundError(errors.New("data not found").Error()))
	}

	return MapPatientResponse(patient)
}
