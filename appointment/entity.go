package appointment

import (
	"GetfitWithPhysio-backend/patient"
	"GetfitWithPhysio-backend/service"
	"time"
)

type Appointment struct {
	Id               int
	Id_patient       int
	Id_service       int
	Appointment_date time.Time
	Description      string
	Status           string
	Patient          patient.Patient `gorm:"ForeignKey: Id_patient"`
	Service          service.Service `gorm:"ForeignKey: Id_service"`
}
