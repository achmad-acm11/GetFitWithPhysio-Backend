package patient

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type ReposioryPatient interface {
	GetAll(ctx context.Context, tx *sql.Tx) []Patient
}

type repositoryPatient struct {
}

func NewRepositoryPatient() *repositoryPatient {
	return &repositoryPatient{}
}

func (r *repositoryPatient) GetAll(ctx context.Context, tx *sql.Tx) []Patient {
	query := "SELECT * FROM patients"

	data, err := tx.QueryContext(ctx, query)
	helper.HandleError(err)
	defer data.Close()

	var patients []Patient
	for data.Next() {
		patient := Patient{}

		data.Scan(&patient.Id, &patient.Id_user, &patient.Gender, &patient.Nik, &patient.Birth_date, &patient.Phone, &patient.Address, &patient.Occupation)

		patients = append(patients, patient)
	}

	return patients
}
