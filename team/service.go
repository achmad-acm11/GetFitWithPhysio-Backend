package team

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ServiceTeam interface {
	GetAllTeams(ctx context.Context) []TeamResponse
}

type serivceTeam struct {
	repo     RepositoryTeam
	db       *sql.DB
	validate *validator.Validate
}

func NewTeamService(repo RepositoryTeam, db *sql.DB, validate *validator.Validate) *serivceTeam {
	return &serivceTeam{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

// Service Get All Data Teams
func (s *serivceTeam) GetAllTeams(ctx context.Context) []TeamResponse {
	//Start Transaction
	tx, err := s.db.Begin()
	helper.HandleError(err)
	// Handle Transaction
	defer helper.CommitOrRollback(tx)

	teams, err := s.repo.GetAll(ctx, tx)
	helper.HandleError(err)

	return MapTeamsResponse(teams)
}
