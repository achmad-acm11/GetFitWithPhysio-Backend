package patient

import (
	"GetfitWithPhysio-backend/user"
)

type PatientResponse struct {
	Id      int    `json:"id"`
	Id_user int    `json:"id_user"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Photo   string `json:"photo"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type RegisterResponse struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Nik        string `json:"nik"`
	Birth_date string `json:"birthdate"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type CreatePatientResponse struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Nik        string `json:"nik"`
	Birth_date string `json:"birthdate"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type DetailPatientResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Photo_user string `json:"photo"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Nik        string `json:"nik"`
	Birth_date string `json:"birthdate"`
	Email      string `json:"email"`
}

func MapPatientResponse(patient Patient) PatientResponse {
	return PatientResponse{
		Id:      patient.Id,
		Id_user: patient.Id_user,
		Name:    patient.User.Name,
		Gender:  patient.Gender,
		Photo:   patient.User.Photo_user,
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
		Name:       user.Name,
		Gender:     patient.Gender,
		Phone:      patient.Phone,
		Address:    patient.Address,
		Nik:        patient.Nik,
		Birth_date: patient.Birth_date.Format("02-01-2006"),
		Email:      user.Email,
		Password:   user.Password,
	}
}
func MapCreatePatientResponse(patient Patient, user user.User) CreatePatientResponse {

	return CreatePatientResponse{
		Name:       user.Name,
		Gender:     patient.Gender,
		Phone:      patient.Phone,
		Address:    patient.Address,
		Nik:        patient.Nik,
		Birth_date: patient.Birth_date.Format("02-01-2006"),
		Email:      user.Email,
		Password:   user.Password,
	}
}
func MapDetailPatientResponse(patient Patient) DetailPatientResponse {

	return DetailPatientResponse{
		Id:         patient.Id,
		Name:       patient.User.Name,
		Gender:     patient.Gender,
		Photo_user: patient.User.Photo_user,
		Phone:      patient.Phone,
		Address:    patient.Address,
		Nik:        patient.Nik,
		Birth_date: patient.Birth_date.Format("02-01-2006"),
		Email:      patient.User.Email,
	}
}
