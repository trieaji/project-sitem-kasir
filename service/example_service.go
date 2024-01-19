package service

import (
	"context"
	"projectsistemkasir/model/web"
)

type ExampleService interface {
	Create(ctx context.Context, request web.ExampleCreateRequest) web.ExampleResponse
}