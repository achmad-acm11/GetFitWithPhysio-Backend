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

func (c *controllerPatient) Register(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	registerRequest := RegisterRequest{}

	helper.ReadFromRequestBody(req, &registerRequest)

	data := c.service.Register(req.Context(), registerRequest)

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Register Success",
			Status:  "success",
			Code:    200,
		},
		Data: data,
	}
	helper.WriteToResponsebody(res, response)
}
func (c *controllerPatient) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	createPatientRequest := CreatePatientRequest{}

	helper.ReadFromRequestBody(req, &createPatientRequest)

	data := c.service.CreateService(req.Context(), createPatientRequest)

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Create Success",
			Status:  "success",
			Code:    200,
		},
		Data: data,
	}
	helper.WriteToResponsebody(res, response)
}
