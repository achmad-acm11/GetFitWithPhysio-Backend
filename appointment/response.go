package appointment

type AppointmentResponse struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Appointment_date string `json:"appointment_date"`
	Product          string `json:"product"`
	Gender           string `json:"gender"`
	Address          string `json:"address"`
	Photo            string `json:"photo"`
	Phone            string `json:"phone"`
	Description      string `json:"description"`
	Status           string `json:"status"`
}

func MapAppointmentResponse(appointment Appointment) AppointmentResponse {
	return AppointmentResponse{
		Id:               appointment.Id,
		Name:             appointment.Patient.User.Name,
		Appointment_date: appointment.Appointment_date.Format("02-01-2006"),
		Product:          appointment.Service.Service_name,
		Gender:           appointment.Patient.Gender,
		Address:          appointment.Patient.Address,
		Photo:            appointment.Patient.User.Photo_user,
		Phone:            appointment.Patient.Phone,
		Description:      appointment.Description,
		Status:           appointment.Status,
	}
}

func MapAppointmentsResponse(appointments []Appointment) []AppointmentResponse {
	appointmentsRes := []AppointmentResponse{}

	for _, v := range appointments {
		appointmentsRes = append(appointmentsRes, MapAppointmentResponse(v))
	}

	return appointmentsRes
}
