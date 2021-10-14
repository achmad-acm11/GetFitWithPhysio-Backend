package user

import (
	"GetfitWithPhysio-backend/helper"
	"context"
	"database/sql"
)

type RepositoryUser interface {
	Create(ctx context.Context, tx *sql.Tx, user User) User
}

type repositoryUser struct {
}

func NewRepositoryUser() *repositoryUser {
	return &repositoryUser{}
}

func (r *repositoryUser) Create(ctx context.Context, tx *sql.Tx, user User) User {
	query := "INSERT INTO users (role,name,email,password) VALUES (?,?,?,?)"

	result, err := tx.ExecContext(ctx, query, user.Role, user.Name, user.Email, user.Password)
	helper.HandleError(err)

	userId, err := result.LastInsertId()
	helper.HandleError(err)

	user.Id = int(userId)
	return user
}
