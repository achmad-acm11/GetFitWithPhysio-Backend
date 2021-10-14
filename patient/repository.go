package patient

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type ReposioryPatient interface {
	GetAll(ctx context.Context, tx *sql.Tx) []Patient
	Create(ctx context.Context, tx *sql.Tx, patient Patient) Patient
}

type repositoryPatient struct {
}

func NewRepositoryPatient() *repositoryPatient {
	return &repositoryPatient{}
}

func (r *repositoryPatient) GetAll(ctx context.Context, tx *sql.Tx) []Patient {
	query := "SELECT id,id_user,gender,nik,phone,address,occupation FROM patients"

	data, err := tx.QueryContext(ctx, query)
	helper.HandleError(err)
	defer data.Close()

	var patients []Patient
	for data.Next() {
		patient := Patient{}

		// s := reflect.ValueOf(&patient).Elem()
		// numCols := s.NumField()
		// columns := make([]interface{}, numCols)

		// for i := 0; i < numCols; i++ {
		// 	field := s.Field(i)
		// 	columns[i] = field.Addr().Interface()
		// }
		// fmt.Println(columns...)
		// data.Scan(columns...)

		data.Scan(&patient.Id, &patient.Id_user, &patient.Gender, &patient.Nik, &patient.Phone, &patient.Address, &patient.Occupation)
		patients = append(patients, patient)
	}

	return patients
}

func (r *repositoryPatient) Create(ctx context.Context, tx *sql.Tx, patient Patient) Patient {
	query := "INSERT INTO patients (id_user,gender,nik,phone,address) VALUES (?,?,?,?,?)"

	_, err := tx.ExecContext(ctx, query, patient.Id_user, patient.Gender, patient.Nik, patient.Phone, patient.Address)
	helper.HandleError(err)

	return patient
}
