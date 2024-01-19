package repository

import (
	"context"
	"database/sql"
	"projectsistemkasir/model/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, payment domain.Payments) domain.Payments
	Update(ctx context.Context, tx *sql.Tx, payment domain.Payments) domain.Payments
	Delete(ctx context.Context, tx *sql.Tx, payment domain.Payments)
	FindById(ctx context.Context, tx *sql.Tx, paymentId int) (domain.Payments, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Payments
}