package repository

import (
	"context"
	"errors"
	"database/sql"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (repository *PaymentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, payment domain.Payments) domain.Payments {
	SQL := "INSERT INTO payments(name, type) VALUES (?,?)"
	result, err := tx.ExecContext(ctx, SQL, payment.Name, payment.Type)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	payment.Id = int(id)
	return payment
}

func (repository *PaymentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, payment domain.Payments) domain.Payments {
	SQL := "update payments set name = ?, type = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, payment.Name, payment.Type, payment.Id)
	helper.PanicIfError(err)

	return payment
}

func (repository *PaymentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, payment domain.Payments) {
	SQL := "delete from payments where id = ?"
	_, err := tx.ExecContext(ctx, SQL, payment.Id)
	helper.PanicIfError(err)
}

func (repository *PaymentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, paymentId int) (domain.Payments, error) {
	SQL := "select id, name, type from payments where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, paymentId)
	helper.PanicIfError(err)
	defer rows.Close()

	payment := domain.Payments{}
	if rows.Next() {
		err := rows.Scan(&payment.Id, &payment.Name, &payment.Type)
		helper.PanicIfError(err)
		return payment, nil
	} else {
		return payment, errors.New("payment is not found")
	}
}

func (repository *PaymentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Payments {
	SQL := "select id, name, type from payments"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var payments []domain.Payments
	for rows.Next() {
		payment := domain.Payments{}
		err := rows.Scan(&payment.Id, &payment.Name, &payment.Type)
		helper.PanicIfError(err)
		payments = append(payments, payment)
	}

	return payments
}