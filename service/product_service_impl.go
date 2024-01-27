package service

import (
	"context"
	"database/sql"
	// "fmt"
	"projectsistemkasir/exception"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
	"projectsistemkasir/model/web"
	"projectsistemkasir/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	UserRepository repository.UserRepository
	ProductRepository repository.ProductRepository
	DB *sql.DB
	Validate validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB: DB,
		Validate: *validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Products {
		Sku: request.Sku,
		Name: request.Name,
		Stock: request.Stock,
		Price: request.Price,
		CategoryId: request.CategoryId,
	}

	product = service.ProductRepository.Save(ctx,tx,product)

	return helper.ToProductResponse(product)
	
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	//saya mendapatkan data user nya terlebih dahulu
	// userRole, err := service.UserRepository.FindById(ctx,tx,request.Id)
	// if err != nil {
	// 	panic(exception.NewNotFoundError(err.Error()))
	// }
	// fmt.Println("=== userRole ===")
	// fmt.Println(userRole)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Sku = request.Sku
	product.Name = request.Name
	product.Stock = request.Stock
	product.Price = request.Price
	product.CategoryId = request.CategoryId

	product = service.ProductRepository.Update(ctx,tx,product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, user)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	return helper.ToProductResponses(products)
}