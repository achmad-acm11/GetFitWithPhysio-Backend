package service

type ServiceResponse struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Service_name string `json:"service_name"`
	Description  string `json:"description"`
	Price        int    `json:"price"`
	Discount     int    `json:"discount"`
}

func MapServiceResponse(service Service) ServiceResponse {
	return ServiceResponse{
		Id:           service.Id,
		Image:        service.Image,
		Service_name: service.Service_name,
		Description:  service.Description,
		Price:        service.Price,
		Discount:     service.Promo.Discount,
	}
}
func MapServicesResponse(services []Service) []ServiceResponse {
	var servicesReponse []ServiceResponse

	for _, v := range services {
		servicesReponse = append(servicesReponse, MapServiceResponse(v))
	}

	return servicesReponse
}
