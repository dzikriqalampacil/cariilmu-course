package repository

import (
	"context"
	"database/sql"

	"github.com/dzikriqalampacil/cariilmu-course/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
	FindById(ctx context.Context, tx *sql.Tx, userId int64) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
}
