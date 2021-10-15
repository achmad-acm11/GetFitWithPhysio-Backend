package service

import "GetfitWithPhysio-backend/promo"

type Service struct {
	Id           int
	Kode_promo   int
	Service_name string
	Kuota_meet   int
	Price        int
	Image        string
	Description  string
	Promo        promo.Promo `gorm:"ForeignKey:Kode_promo"`
}
