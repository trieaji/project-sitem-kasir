package repository

import (
	"context"
	"database/sql"
	"projectsistemkasir/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	Update(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	Delete(ctx context.Context, tx *sql.Tx, user domain.Users)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.Users, error)
	FindEmail(ctx context.Context, tx *sql.Tx, userMail any) (domain.Users, error)
	UpdateTkn(ctx context.Context, tx *sql.Tx, user domain.Users) domain.Users
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Users
}