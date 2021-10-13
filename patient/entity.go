package patient

import "time"

type Patient struct {
	Id         int
	Id_user    int
	Gender     string
	Nik        string
	Birth_date time.Time
	Phone      string
	Address    string
	Occupation string
}
