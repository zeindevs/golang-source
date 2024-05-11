package service

import (
	"errors"
	"ginjwt/config"
	"ginjwt/data/request"
	"ginjwt/helper"
	"ginjwt/model"
	"ginjwt/repository"
	"ginjwt/utils"

	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UsersRepository
	Validate       *validator.Validate
}

func NewAuthenticationServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

// Login implements AuthenticationService.
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	// find username in database
	new_user, user_err := a.UserRepository.FindByUsername(users.Username)
	if user_err != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_user.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate Token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil
}

// Register implements AuthenticationService.
func (a *AuthenticationServiceImpl) Register(users request.CreateUserRequest) {
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUser := model.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}

	a.UserRepository.Save(newUser)
}
