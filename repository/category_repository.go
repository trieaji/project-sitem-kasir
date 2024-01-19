package repository

import (
	"context"
	"database/sql"
	"projectsistemkasir/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Categories) domain.Categories
	Update(ctx context.Context, tx *sql.Tx, category domain.Categories) domain.Categories
	Delete(ctx context.Context, tx *sql.Tx, category domain.Categories)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Categories, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Categories
}