package team

type CreateTeamRequest struct {
	Name        string `validate:"required" form:"name"`
	Position    string `validate:"required" form:"position"`
	Url         string `validate:"required" form:"url"`
	Description string `validate:"required" form:"description"`
}
