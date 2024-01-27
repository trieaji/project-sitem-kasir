package repository

import (
	"context"
	"errors"
	"database/sql"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "INSERT INTO products(sku, name, stock, price, category_id) VALUES (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, product.Sku ,product.Name, product.Stock, product.Price, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Products) domain.Products {
	SQL := "update products set sku = ?, name = ?, stock = ?, price = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Sku, product.Name, product.Stock, product.Price, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Products) {
	SQL := "delete from products where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Products, error) {
	SQL := "select id, sku ,name, stock, price, category_id from products where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Products{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Sku, &product.Name, &product.Stock, &product.Price, &product.CategoryId)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Products {
	SQL := "select id, sku, name, stock, price, category_id from products"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Products
	for rows.Next() {
		product := domain.Products{}
		err := rows.Scan(&product.Id, &product.Sku, &product.Name, &product.Stock, &product.Price, &product.CategoryId)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return products
}