package team

import (
	"GetfitWithPhysio-backend/helper"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-playground/form"
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

func (c *controllerTeam) CreateTeam(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// Create Team Start
	createTeamRequest := CreateTeamRequest{}

	values := url.Values{}
	values.Add("name", req.FormValue("name"))
	values.Add("position", req.FormValue("position"))
	values.Add("url", req.FormValue("url"))
	values.Add("description", req.FormValue("description"))

	form.NewDecoder().Decode(&createTeamRequest, values)

	data := c.service.Create(req.Context(), createTeamRequest)
	// Create Team End
	if _, _, err := req.FormFile("photo"); err == nil {
		// Save Upload File
		pathFile := helper.UploadPhoto(req, req.FormValue("name")+strconv.Itoa(data.Id), "photo", "resources/team_photos")
		// Update Photo
		c.service.UploadPhoto(req.Context(), data.Id, pathFile)

		data.Photo_team = pathFile
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
