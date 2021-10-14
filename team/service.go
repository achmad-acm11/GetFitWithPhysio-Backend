package team

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceTeam interface {
	GetAllTeams(ctx context.Context) []TeamResponse
}

type serivceTeam struct {
	repo     RepositoryTeam
	db       *gorm.DB
	validate *validator.Validate
}

func NewTeamService(repo RepositoryTeam, db *gorm.DB, validate *validator.Validate) *serivceTeam {
	return &serivceTeam{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}

// Service Get All Data Teams
func (s *serivceTeam) GetAllTeams(ctx context.Context) []TeamResponse {
	//Start Transaction
	tx := s.db.Begin()
	// Handle Transaction
	defer helper.CommitOrRollback(tx)

	teams := s.repo.GetAll(ctx, tx)

	return MapTeamsResponse(teams)
}
