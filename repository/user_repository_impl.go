package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dzikriqalampacil/cariilmu-course/helper"
	"github.com/dzikriqalampacil/cariilmu-course/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, User domain.User) domain.User {
	SQL := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id"
	var id int64
	err := tx.QueryRowContext(ctx, SQL, User.Name, User.Email, User.Password).Scan(&id)
	helper.PanicIfError(err)

	User.Id = id
	return User
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, User domain.User) domain.User {
	SQL := "UPDATE users SET name = $2, email = $3, password = $4  WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, User.Id, User.Name, User.Email, User.Password)
	helper.PanicIfError(err)

	return User
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, User domain.User) {
	SQL := "DELETE FROM users WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, User.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, UserId int64) (domain.User, error) {
	SQL := "SELECT id, name, email, password FROM users WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, UserId)
	helper.PanicIfError(err)
	defer rows.Close()

	User := domain.User{}
	if rows.Next() {
		err := rows.Scan(&User.Id, &User.Name, &User.Email, &User.Password)
		helper.PanicIfError(err)
		return User, nil
	} else {
		return User, errors.New("User is not found")
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, name, email, password FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var Users []domain.User
	for rows.Next() {
		User := domain.User{}
		err := rows.Scan(&User.Id, &User.Name, &User.Email, &User.Password)
		helper.PanicIfError(err)
		Users = append(Users, User)
	}
	return Users
}
