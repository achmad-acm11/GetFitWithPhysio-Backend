package transaction

import (
	"GetfitWithPhysio-backend/patient"
	"GetfitWithPhysio-backend/service"
	"GetfitWithPhysio-backend/user"
	"encoding/json"
	"strconv"
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
	Patient    patient.Patient `gorm:"Foreignkey:Id_user"`
}

type TmpTransactions struct {
	Id           int    `json:"id"`
	Id_user      int    `json:"id_user"`
	Id_service   int    `json:"id_service"`
	Amount       int    `json:"amount"`
	Code         string `json:"code"`
	Status       string `json:"status"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Photo_user   string `json:"photo_user"`
	Occupation   string `json:"occupation"`
	Nik          string `json:"nik"`
	Gender       string `json:"gender"`
	Phone        int    `json:"phone"`
	Address      string `json:"address"`
	Kode_promo   int    `json:"kode_promo"`
	Service_name string `json:"service_name"`
	Kuota_meet   int    `json:"kuota_meet"`
	Price        int    `json:"price"`
	Image        string `json:"image"`
	Description  string `json:"description"`
}

func MapTransaction(result map[string]interface{}) Transaction {
	byte, _ := json.Marshal(result)
	tmp := TmpTransactions{}
	json.Unmarshal(byte, &tmp)

	return Transaction{
		Id:         tmp.Id,
		Id_user:    tmp.Id_user,
		Id_service: tmp.Id_service,
		Amount:     tmp.Amount,
		Code:       tmp.Code,
		Status:     tmp.Status,
		User: user.User{
			Name:       tmp.Name,
			Email:      tmp.Email,
			Password:   tmp.Password,
			Photo_user: tmp.Photo_user,
		},
		Patient: patient.Patient{
			Occupation: tmp.Occupation,
			Nik:        tmp.Nik,
			Gender:     tmp.Gender,
			Phone:      strconv.Itoa(tmp.Phone),
			Address:    tmp.Address,
		},
		Service: service.Service{
			Kode_promo:   tmp.Kode_promo,
			Service_name: tmp.Service_name,
			Kuota_meet:   tmp.Kuota_meet,
			Price:        tmp.Price,
			Image:        tmp.Image,
			Description:  tmp.Description,
		},
	}
}
func MapTransactions(result []map[string]interface{}) []Transaction {
	transactions := []Transaction{}

	for _, v := range result {
		transactions = append(transactions, MapTransaction(v))
	}
	return transactions
}
