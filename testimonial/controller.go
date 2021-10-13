package testimonial

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type controllerTestimonial struct {
	service ServiceTestimonial
}

func NewControllrerTestimnoial(service ServiceTestimonial) *controllerTestimonial {
	return &controllerTestimonial{
		service: service,
	}
}

func (c *controllerTestimonial) GetAllController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
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
