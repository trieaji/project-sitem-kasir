package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"projectsistemkasir/exception"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
	"projectsistemkasir/model/web"
	"projectsistemkasir/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"

	// "golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB *sql.DB
	Validate validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: *validate,
	}
}


func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pass := request.Password
	// bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	hash, _ := helper.HashPassword(pass)

	user := domain.Users {
		Name: request.Name,
		Email: request.Email,
		Password: string(hash),
		Role: request.Role,
	}

	user = service.UserRepository.Save(ctx,tx,user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginUser) web.UserLoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindEmail(ctx, tx, request.Email)
	fmt.Println("===user===")
	fmt.Println(user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var uMail = user.Email
	fmt.Println("===uMail===")
	fmt.Println(uMail)
	if length := len(uMail); length == 0 {
		fmt.Println("Username or password wrong")
	}
	fmt.Println("===umail2===")
	fmt.Println(uMail)

	pass := request.Password
	fmt.Println("===pass===")
	fmt.Println(pass)
	hash, _ := helper.HashPassword(pass)
	fmt.Println("===hash===")
	fmt.Println(hash)
	isPasswordValid := helper.CheckPasswordHash(pass, hash)
	fmt.Println("===isPasswordValid===")
	fmt.Println(isPasswordValid)
	if !isPasswordValid {
		fmt.Println("Username or password wrong")
	}
	
	//saatnya untuk generate tokennya
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := helper.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
	}
	fmt.Println("===token===")
	fmt.Println(token)

	// loginResponse := web.UserLoginResponse {
	// 	Token: token,
	// }
	

	user.Token = token
	fmt.Println("===usertoken===")
	fmt.Println(user)

	user = service.UserRepository.UpdateTkn(ctx,tx,user)
	fmt.Println("===userupdate===")
	fmt.Println(user)

	return helper.ToLoginResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	pass := request.Password
	hash, _ := helper.HashPassword(pass)

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hash)
	user.Role = request.Role

	user = service.UserRepository.Update(ctx,tx,user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}


func (service *UserServiceImpl) Logut(ctx context.Context, request web.LogoutUser) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Token = "null"

	service.UserRepository.UpdateTkn(ctx,tx,user)

}