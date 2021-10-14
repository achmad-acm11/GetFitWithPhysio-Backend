package team

type TeamResponse struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Description_profile string `json:"description"`
	Position            string `json:"position"`
	Url                 string `json:"url"`
	Photo_team          string `json:"photo_team"`
}

func MapTeamResponse(team Team) TeamResponse {
	return TeamResponse{
		Id:                  team.Id,
		Name:                team.Name,
		Description_profile: team.Description_profile,
		Position:            team.Position,
		Url:                 team.Url,
		Photo_team:          team.Photo_team,
	}
}
func MapTeamsResponse(teams []Team) []TeamResponse {
	var teamsResponse []TeamResponse
	for _, v := range teams {
		teamsResponse = append(teamsResponse, MapTeamResponse(v))
	}
	return teamsResponse
}
