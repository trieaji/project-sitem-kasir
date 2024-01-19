package repository

import (
	"context"
	"errors"
	"database/sql"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "INSERT INTO users(name, email, password, role) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Role)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "update users set name = ?, email = ?, password = ?, role = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Role, user.Id)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.Users) {
	SQL := "delete from users where id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.Users, error) {
	SQL := "select id, name, email, password, role from users where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.Users{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

func (repository *UserRepositoryImpl) FindEmail(ctx context.Context, tx *sql.Tx, userMail any) (domain.Users, error) {
	SQL := "select email, password from users where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, userMail)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.Users{}
	if rows.Next() {
		err := rows.Scan(&user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("email is not found")
	}
}

func (repository *UserRepositoryImpl) UpdateTkn(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users {
	SQL := "update users set token = ? where email = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Token, user.Email)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Users {
	SQL := "select id, name, email, password, role from users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.Users
	for rows.Next() {
		user := domain.Users{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}