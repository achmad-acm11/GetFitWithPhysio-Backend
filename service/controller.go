package service

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-playground/form"
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

// Get All Service Controller
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

func (c *controllerService) CreateController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	requestService := CreateServiceRequest{}

	values := url.Values{}
	values.Add("service_name", req.FormValue("service_name"))
	values.Add("kuota_meet", req.FormValue("kuota_meet"))
	values.Add("price", req.FormValue("price"))
	values.Add("description", req.FormValue("description"))
	values.Add("kode_promo", req.FormValue("kode_promo"))

	form.NewDecoder().Decode(&requestService, values)

	data := c.service.CreateService(req.Context(), requestService)

	if _, _, err := req.FormFile("image"); err == nil {

		filePath := helper.UploadPhoto(req, req.FormValue("service_name")+strconv.Itoa(data.Id), "image", "resources/service_image", "service_image")

		c.service.UploadImageService(req.Context(), data.Id, filePath)

		data.Image = filePath
	}

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
