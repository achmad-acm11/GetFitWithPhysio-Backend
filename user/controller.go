package user

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type controllerUser struct {
	service ServiceUser
}

func NewControllerUser(service ServiceUser) *controllerUser {
	return &controllerUser{
		service: service,
	}
}

func (c *controllerUser) Login(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	loginRequest := LoginRequest{}

	helper.ReadFromRequestBody(req, &loginRequest)

	c.service.Login(req.Context(), loginRequest)

	response := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Login Success",
			Status:  "success",
			Code:    200,
		},
		Data: nil,
	}

	helper.WriteToResponsebody(res, response)
}
