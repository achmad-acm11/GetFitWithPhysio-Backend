package patient

import (
	"GetfitWithPhysio-backend/user"
)

type PatientResponse struct {
	Id      int    `json:"id"`
	Id_user int    `json:"id_user"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type RegisterResponse struct {
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Nik      string `json:"nik"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func MapPatientResponse(patient Patient) PatientResponse {
	return PatientResponse{
		Id:      patient.Id,
		Id_user: patient.Id_user,
		Gender:  patient.Gender,
		Phone:   patient.Phone,
		Address: patient.Address,
	}
}

func MapPatientsResponse(patients []Patient) []PatientResponse {
	var patientsRes []PatientResponse

	for _, v := range patients {
		patientsRes = append(patientsRes, MapPatientResponse(v))
	}
	return patientsRes
}

func MapRegisterResponse(patient Patient, user user.User) RegisterResponse {
	return RegisterResponse{
		Name:     user.Name,
		Gender:   patient.Gender,
		Phone:    patient.Phone,
		Address:  patient.Address,
		Nik:      patient.Nik,
		Email:    user.Email,
		Password: user.Password,
	}
}
