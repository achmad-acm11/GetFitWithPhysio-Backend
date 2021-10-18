package appointment

import (
	"GetfitWithPhysio-backend/helper"
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
	repo     RepositoryAppointment
	db       *gorm.DB
	validate *validator.Validate
}

func NewServiceAppoinment(repo RepositoryAppointment, db *gorm.DB, validate *validator.Validate) *serviceAppointment {
	return &serviceAppointment{
		repo:     repo,
		db:       db,
		validate: validate,
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

	date, _ := time.Parse("02-01-2006 15:00:00", reqAppointment.Appointment_date)

	appointment := s.repo.Create(ctx, tx, Appointment{
		Id_patient:       reqAppointment.IdPatient,
		Id_service:       reqAppointment.IdService,
		Appointment_date: date,
		Description:      reqAppointment.Description,
		Status:           "Pending",
	})

	return MapAppointmentResponse(appointment)
}

// Get Detail Data Appointment Service
func (s *serviceAppointment) DetailService(ctx context.Context, appointmentId int) AppointmentResponse {
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	appointment := s.repo.GetOneById(ctx, tx, appointmentId)

	return MapAppointmentResponse(appointment)
}
