package team

import (
	"GetfitWithPhysio-backend/helper"
	"context"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceTeam interface {
	GetAllTeams(ctx context.Context) []TeamResponse
	Create(ctx context.Context, req CreateTeamRequest) TeamResponse
	UploadPhoto(ctx context.Context, teamId int, filePath string)
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

// Service Create Team
func (s *serivceTeam) Create(ctx context.Context, req CreateTeamRequest) TeamResponse {

	// Validate
	err := s.validate.Struct(req)
	helper.HandleError(err)

	// Start Transaction
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	team := s.repo.Create(ctx, tx, Team{
		Name:                req.Name,
		Position:            req.Position,
		Url:                 req.Url,
		Description_profile: req.Description,
	})

	return MapTeamResponse(team)
}

// Service Upload Photo
func (s *serivceTeam) UploadPhoto(ctx context.Context, teamId int, filePath string) {
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	s.repo.UpdatePhoto(ctx, tx, teamId, filePath)

}
