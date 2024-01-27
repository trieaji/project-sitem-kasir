package helper

import (
	"projectsistemkasir/model/domain"
	"projectsistemkasir/model/web"
)

func ToCategoryResponse(category domain.Categories) web.CategoryResponse {
	return web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Categories) []web.CategoryResponse {
	var categoryResponse []web.CategoryResponse

	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}

	return categoryResponse
}

func ToExampleResponse(example domain.Examples) web.ExampleResponse {
	return web.ExampleResponse{
		Id: example.Id,
		Name: example.Name,
		Jurusan: example.Jurusan,
	}
}

func ToPaymentResponse(payment domain.Payments) web.PaymentResponse {
	return web.PaymentResponse{
		Id: payment.Id,
		Name: payment.Name,
		Type: payment.Type,
	}
}

func ToPaymentResponses(payments []domain.Payments) []web.PaymentResponse {
	var paymentResponse []web.PaymentResponse

	for _, payment := range payments {
		paymentResponse = append(paymentResponse, ToPaymentResponse(payment))
	}

	return paymentResponse
}

func ToUserResponse(user domain.Users) web.UserResponse {
	return web.UserResponse{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
	}
}

func ToLoginResponse(user domain.Users) web.UserLoginResponse {
	return web.UserLoginResponse{
		Token: user.Token,
		// Role: user.Role,
	}
}

func ToUserResponses(users []domain.Users) []web.UserResponse {
	var userResponse []web.UserResponse

	for _, user := range users {
		userResponse = append(userResponse, ToUserResponse(user))
	}

	return userResponse
}

func ToProductResponse(product domain.Products) web.ProductResponse {
	return web.ProductResponse {
		Id: product.Id,
		Sku: product.Sku,
		Name: product.Name,
		Stock: product.Stock,
		Price: product.Price,
		CategoryId: product.CategoryId,
	}
}

func ToProductResponses(products []domain.Products) []web.ProductResponse {
	var productResponse []web.ProductResponse

	for _, product := range products {
		productResponse = append(productResponse, ToProductResponse(product))
	}

	return productResponse
}