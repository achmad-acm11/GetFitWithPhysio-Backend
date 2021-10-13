package promo

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type controllerPromo struct {
	service ServicePromo
}

func NewControllerPromo(service ServicePromo) *controllerPromo {
	return &controllerPromo{
		service: service,
	}
}

func (c *controllerPromo) GetAllContoller(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	promosResponse := c.service.GetAllService(req.Context())

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Get Success",
			Status:  "success",
			Code:    200,
		},
		Data: promosResponse,
	}

	helper.WriteToResponsebody(res, response)
}
