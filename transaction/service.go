package transaction

import (
	"GetfitWithPhysio-backend/exception"
	"GetfitWithPhysio-backend/helper"
	"GetfitWithPhysio-backend/service"
	"GetfitWithPhysio-backend/user"
	"context"
	"strconv"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type ServiceTransaction interface {
	GetAllService(ctx context.Context) []TransactionResponse
	CreateService(ctx context.Context, req RequestTransaction) TransactionResponse
}

type serviceTransaction struct {
	repo        RepositoryTransaction
	repoService service.RepositoryService
	repoUser    user.RepositoryUser
	db          *gorm.DB
	validator   *validator.Validate
}

func NewServiceTransaction(repo RepositoryTransaction, repoService service.RepositoryService, repoUser user.RepositoryUser, db *gorm.DB, validator *validator.Validate) *serviceTransaction {
	return &serviceTransaction{
		repo:        repo,
		repoService: repoService,
		repoUser:    repoUser,
		db:          db,
		validator:   validator,
	}
}

// Get All Data Transactions
func (s *serviceTransaction) GetAllService(ctx context.Context) []TransactionResponse {
	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)

	// Get All Data
	transactions := s.repo.GetAll(ctx, tx)

	return MapTransactionsResponse(transactions)
}

// Create Transaction Service
func (s *serviceTransaction) CreateService(ctx context.Context, req RequestTransaction) TransactionResponse {

	tx := s.db.Begin()
	defer helper.CommitOrRollback(tx)
	// Get Data User
	user := s.repoUser.GetUserById(ctx, tx, req.IdUser)

	// Check
	if user.Id == 0 {
		panic(exception.NewNotFoundError("User Not Found"))
	}

	// Get One Service
	service := s.repoService.GetOneById(ctx, tx, req.IdService)

	// Check
	if service.Id == 0 {
		panic(exception.NewNotFoundError("Product Not Found"))
	}

	// Create Transaction
	transaction := s.repo.Create(ctx, tx, Transaction{
		Id_user:    req.IdUser,
		Id_service: req.IdService,
		Amount:     service.Price,
		Status:     "Pending",
	})

	// Generate Code Transaction
	transaction.Code = "A" + strconv.Itoa(transaction.Id)
	// Input Code Transaction
	transaction = s.repo.Update(ctx, tx, transaction)

	transaction.Service.Service_name = service.Service_name
	transaction.User.Name = user.Name

	return MapTransactionResponse(transaction)

}
