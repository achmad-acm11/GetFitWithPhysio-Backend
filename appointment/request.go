package appointment

type AppointmentRequest struct {
	Appointment_date string `validate:"required" json:"appointment_date"`
	Description      string `validate:"required" json:"description"`
	IdPatient        int
	IdService        int
}
