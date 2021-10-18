package patient

import (
	"GetfitWithPhysio-backend/user"
	"time"
)

type Patient struct {
	Id         int
	Id_user    int
	Gender     string
	Nik        string
	Birth_date time.Time // set default NULL if not set value
	Phone      string
	Address    string
	Occupation string
	User       user.User `gorm:"ForeignKey:Id_user"`
}
