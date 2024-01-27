package repository

import (
	"context"
	"database/sql"
	"projectsistemkasir/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	Update(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products
	Delete(ctx context.Context, tx *sql.Tx, product domain.Products)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Products, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Products
}