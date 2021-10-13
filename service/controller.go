package service

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type controllerService struct {
	service ServiceImpl
}

func NewServiceController(service ServiceImpl) *controllerService {
	return &controllerService{
		service: service,
	}
}

func (c *controllerService) GetAllController(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	servicesResponse := c.service.GetAllService(req.Context())

	reponse := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Get Success",
			Status:  "success",
			Code:    200,
		},
		Data: servicesResponse,
	}

	helper.WriteToResponsebody(wr, reponse)
}
