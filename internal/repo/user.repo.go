package repo

import (
	"context"

	"github.com/LeMinh0706/ChatApp/db"
)

type UserRepo struct {
	queries *db.Queries
}

func NewUserRepo(queries *db.Queries) (*UserRepo, error) {
	return &UserRepo{
		queries: queries,
	}, nil
}

func (repo *UserRepo) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return repo.queries.CreateUser(ctx, arg)
}

func (repo *UserRepo) GetUser(ctx context.Context, username string) (db.User, error) {
	return repo.queries.GetUser(ctx, username)
}

func (repo *UserRepo) GetById(ctx context.Context, id int64) (db.GetUserByIdRow, error) {
	return repo.queries.GetUserById(ctx, id)
}
