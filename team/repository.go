package team

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type RepositoryTeam interface {
	GetAll(cx context.Context, tx *sql.Tx) ([]Team, error)
}

type repository struct {
}

// Function for access repo struct
func NewRepository() *repository {
	return &repository{}
}

// SQL Query Get All Data Team
func (r *repository) GetAll(cx context.Context, tx *sql.Tx) ([]Team, error) {
	query := "SELECT * FROM teams"

	// Execute SQL Query
	data, err := tx.QueryContext(cx, query)
	helper.HandleError(err)
	defer data.Close()

	// Mapping Team to Entity
	var teams []Team
	for data.Next() {
		team := Team{}

		err := data.Scan(&team.Id, &team.Name, &team.Position, &team.Url, &team.Description_profile, &team.Photo_team)
		helper.HandleError(err)

		teams = append(teams, team)
	}

	return teams, nil
}
