package patient

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type controllerPatient struct {
	service ServicePatient
}

func NewControllerPatinet(service ServicePatient) *controllerPatient {
	return &controllerPatient{
		service: service,
	}
}

func (c *controllerPatient) GetAllController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	data := c.service.GetAllService(req.Context())

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Get Success",
			Status:  "success",
			Code:    200,
		},
		Data: data,
	}

	helper.WriteToResponsebody(res, response)
}
