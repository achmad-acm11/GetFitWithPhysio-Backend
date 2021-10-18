package galery

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"
	"net/url"

	"github.com/go-playground/form"
	"github.com/julienschmidt/httprouter"
)

type controllerGalery struct {
	service ServiceGalery
}

func NewControllerGalery(service ServiceGalery) *controllerGalery {
	return &controllerGalery{
		service: service,
	}
}

func (c *controllerGalery) GetAllController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
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

func (c *controllerGalery) CreateGaleryController(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	reqCreateGalery := CreateGaleryRequest{}

	values := url.Values{}
	values.Add("caption", req.FormValue("caption"))
	values.Add("sub_caption", req.FormValue("sub_caption"))

	helper.HandleError(form.NewDecoder().Decode(&reqCreateGalery, values))

	data := c.service.CreateService(req.Context(), reqCreateGalery)

	if _, _, err := req.FormFile("photo"); err == nil {
		filePath := helper.UploadPhoto(req, req.FormValue("caption"), "photo", "resources/galeries", "galery")

		c.service.UploadPhoto(req.Context(), data.Id, filePath)

		data.Photo = filePath
	}

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Success Upload Photo",
			Status:  "success",
			Code:    200,
		},
		Data: data,
	}

	helper.WriteToResponsebody(res, response)

}
