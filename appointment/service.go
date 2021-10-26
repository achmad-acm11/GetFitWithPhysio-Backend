package appointment

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/patient"
	"GetfitWithPhysio-backend/service"
	"context"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceAppointment interface {
	GetAllService(ctx context.Context) []AppointmentResponse
	DetailService(cxt context.Context, appointmentId int) AppointmentResponse
	CreateAppointment(ctx context.Context, reqAppointment AppointmentRequest) AppointmentResponse
}

type serviceAppointment struct {
	repo        RepositoryAppointment
	repoPatient patient.ReposioryPatient
	repoService service.RepositoryService
	db          *gorm.DB
	validate    *validator.Validate
}

func NewServiceAppoinment(repo RepositoryAppointment, repoPatient patient.ReposioryPatient, repoService service.RepositoryService, db *gorm.DB, validate *validator.Validate) *serviceAppointment {
	return &serviceAppointment{
		repo:        repo,
		repoPatient: repoPatient,
		repoService: repoService,
		db:          db,
		validate:    validate,
	}
}

// Get All Data Appointment Service
func (s *serviceAppointment) GetAllService(ctx context.Context) []AppointmentResponse {
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	appointments := s.repo.GetAll(ctx, tx)

	return MapAppointmentsResponse(appointments)
}

// Create Data Appointment
func (s *serviceAppointment) CreateAppointment(ctx context.Context, reqAppointment AppointmentRequest) AppointmentResponse {
	err := s.validate.Struct(reqAppointment)
	helper.HandleError(err)

	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	patient := s.repoPatient.GetOneById(ctx, tx, reqAppointment.IdPatient)

	if patient.Id == 0 {
		panic(exception.NewNotFoundError("Patient Not Found"))
	}

	service := s.repoService.GetOneById(ctx, tx, reqAppointment.IdService)

	if service.Id == 0 {
		panic(exception.NewNotFoundError("Service Not Found"))
	}

	date, _ := time.Parse("02-01-2006 15:00:00", reqAppointment.Appointment_date)

	appointment := s.repo.Create(ctx, tx, Appointment{
		Id_patient:       reqAppointment.IdPatient,
		Id_service:       reqAppointment.IdService,
		Appointment_date: date,
		Description:      reqAppointment.Description,
		Status:           "Pending",
	})

	appointment.Patient.User.Name = patient.User.Name
	appointment.Service.Service_name = service.Service_name
	appointment.Patient.Gender = patient.Gender
	appointment.Patient.Address = patient.Address
	appointment.Patient.User.Photo_user = patient.User.Photo_user
	appointment.Patient.Phone = patient.Phone

	return MapAppointmentResponse(appointment)
}

// Get Detail Data Appointment Service
func (s *serviceAppointment) DetailService(ctx context.Context, appointmentId int) AppointmentResponse {
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	appointment := s.repo.GetOneById(ctx, tx, appointmentId)

	if appointment.Id == 0 {
		panic(exception.NewNotFoundError("Appointment Not Found"))
	}
	return MapAppointmentResponse(appointment)
}
