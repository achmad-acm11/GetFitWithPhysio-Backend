package patient

type PatientResponse struct {
	Id      int    `json:"id"`
	Id_user int    `json:"id_user"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
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
