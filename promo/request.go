package promo

type CreatePromoRequest struct {
	Discount int `validate:"required" json:"discount"`
}
