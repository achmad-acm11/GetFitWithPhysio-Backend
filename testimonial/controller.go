package testimonial

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type controllerTestimonial struct {
	service ServiceTestimonial
}

func NewControllrerTestimonial(service ServiceTestimonial) *controllerTestimonial {
	return &controllerTestimonial{
		service: service,
	}
}

// Get All Data Testimonial Controller
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

// Create Data Testimonial Controller
func (c *controllerTestimonial) CreateTestimonialController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	reqTestimonial := RequestTestimonial{}

	helper.ReadFromRequestBody(req, &reqTestimonial)
	reqTestimonial.Id_user = 10

	data := c.service.CreateTestimonialService(req.Context(), reqTestimonial)

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
