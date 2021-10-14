package team

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type controllerTeam struct {
	service ServiceTeam
}

func NewTeamController(service ServiceTeam) *controllerTeam {
	return &controllerTeam{service: service}
}

// Controller endpoint GET "/api/v1/teams"
func (c *controllerTeam) GetAllController(wr http.ResponseWriter, req *http.Request, params httprouter.Params) {
	teamsResponse := c.service.GetAllTeams(req.Context())

	webResponse := helper.FormatResponse{
		Meta: helper.Meta{
			Message: "Get Success",
			Status:  "success",
			Code:    200,
		},
		Data: teamsResponse,
	}

	helper.WriteToResponsebody(wr, webResponse)
}
