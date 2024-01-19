package service

import (
	"context"
	"projectsistemkasir/model/web"
)

type PaymentService interface {
	Create(ctx context.Context, request web.PaymentCreateRequest) web.PaymentResponse
	Update(ctx context.Context, request web.PaymentUpdateRequest) web.PaymentResponse
	Delete(ctx context.Context, paymentId int)
	FindById(ctx context.Context, paymentId int) web.PaymentResponse
	FindAll(ctx context.Context) []web.PaymentResponse
}