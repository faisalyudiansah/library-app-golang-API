package services

import (
	"context"
	"time"

	apperrors "library-api/appErrors"
	"library-api/dtos"
	"library-api/helpers"
	"library-api/models"
	"library-api/repositories"
)

type UserService interface {
	PostRegisterUserService(context.Context, dtos.RequestRegisterUser) (*dtos.ResponseDataUser, error)
	PostLoginUserService(context.Context, dtos.RequestLoginUser) (*dtos.ResponseDataUser, error)
}

type UserServiceImplementation struct {
	UserRepository repositories.UserRepository
	Bcrypt         helpers.Bcrypt
	Jwt            helpers.JWTProvider
}

func NewUserServiceImplementation(us repositories.UserRepository, bc helpers.Bcrypt, jwt helpers.JWTProvider) *UserServiceImplementation {
	return &UserServiceImplementation{
		UserRepository: us,
		Bcrypt:         bc,
		Jwt:            jwt,
	}
}

func (us *UserServiceImplementation) PostRegisterUserService(ctx context.Context, reqBody dtos.RequestRegisterUser) (*dtos.ResponseDataUser, error) {
	if IsEmailAlreadyRegistered := us.UserRepository.IsEmailAlreadyRegistered(ctx, reqBody.Email); IsEmailAlreadyRegistered {
		return nil, apperrors.ErrUserEmailAlreadyExists
	}
	hashPwd, err := us.Bcrypt.HashPassword(reqBody.Password, 10)
	if err != nil {
		return nil, apperrors.ErrUserFailedRegister
	}
	_, err = us.UserRepository.PostUser(ctx, reqBody, string(hashPwd))
	if err != nil {
		return nil, apperrors.ErrUserFailedRegister
	}
	result := dtos.ToDtoResponseUserInfo(&models.User{
		Name:      reqBody.Name,
		Email:     reqBody.Email,
		CreatedAt: time.Now(),
	})
	return &result, nil
}

func (us *UserServiceImplementation) PostLoginUserService(ctx context.Context, reqBody dtos.RequestLoginUser) (*dtos.ResponseDataUser, error) {
	user, err := us.UserRepository.GetUserByEmail(ctx, reqBody.Email)
	if err != nil {
		return nil, apperrors.ErrUserInvalidEmailPassword
	}
	isValid, err := us.Bcrypt.CheckPassword(reqBody.Password, []byte(user.Password))
	if err != nil || !isValid {
		return nil, apperrors.ErrUserInvalidEmailPassword
	}
	accesToken, err := us.Jwt.CreateToken(int64(user.ID))
	if err != nil {
		return nil, apperrors.ErrISE
	}
	result := dtos.ToDtoResponseUserAccessToken(accesToken)
	return &result, nil
}
