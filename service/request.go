package service

type CreateServiceRequest struct {
	Kode_promo   int    `form:"kode_promo"`
	Service_name string `validate:"required" form:"service_name"`
	Kuota_meet   int    `validate:"required" form:"kuota_meet"`
	Price        int    `validate:"required" form:"price"`
	Description  string `validate:"required" form:"description"`
}
