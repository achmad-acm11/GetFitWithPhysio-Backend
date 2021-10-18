package transaction

import (
	"GetfitWithPhysio-backend/service"
	"GetfitWithPhysio-backend/user"
)

type Transaction struct {
	Id         int
	Id_user    int
	Id_service int
	Amount     int
	Code       string
	Status     string
	User       user.User       `gorm:"ForeignKey: Id_user"`
	Service    service.Service `gorm:"ForeignKey: Id_service"`
}
