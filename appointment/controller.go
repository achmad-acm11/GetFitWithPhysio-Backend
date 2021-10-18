package appointment

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type controllerAppointment struct {
	service ServiceAppointment
}

func NewControllerAppointment(service ServiceAppointment) *controllerAppointment {
	return &controllerAppointment{
		service: service,
	}
}

func (c *controllerAppointment) GetAllController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

func (c *controllerAppointment) CreateAppointment(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceId, _ := strconv.Atoi(params.ByName("serviceId"))

	requestAppointment := AppointmentRequest{}

	helper.ReadFromRequestBody(req, &requestAppointment)

	requestAppointment.IdPatient = 1
	requestAppointment.IdService = serviceId

	data := c.service.CreateAppointment(req.Context(), requestAppointment)

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

func (c *controllerAppointment) DetailAppointment(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	appointmentId, _ := strconv.Atoi(params.ByName("appointmentId"))

	data := c.service.DetailService(req.Context(), appointmentId)

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
