package promo

type PromoResponse struct {
	Id       int `json:"id"`
	Discount int `json:"discount"`
}

func MapPromoResponse(promo Promo) PromoResponse {
	return PromoResponse{
		Id:       promo.Id,
		Discount: promo.Discount,
	}
}
func MapPromosResponse(promos []Promo) []PromoResponse {
	var promosResponse []PromoResponse
	for _, v := range promos {
		promosResponse = append(promosResponse, MapPromoResponse(v))
	}
	return promosResponse
}
