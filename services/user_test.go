package services_test

import (
	"context"
	"testing"
	"time"

	apperrors "git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/appErrors"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/dtos"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/models"
	"git.garena.com/sea-labs-id/bootcamp/batch-04/shared-projects/library-api/services"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceImplementation_PostRegisterUserService(t *testing.T) {
	var (
		nameUser     = "Frierennnnnnn"
		emailUser    = "frieren@gmail.com"
		passwordUser = "12345"
		hashPwd      = "54321"
	)
	resUser := &models.User{
		ID:        1,
		Name:      nameUser,
		Email:     emailUser,
		Password:  hashPwd,
		CreatedAt: time.Time{},
		UpdatedAt: &time.Time{},
		DeleteAt:  nil,
	}
	tests := []struct {
		name           string
		reqBody        dtos.RequestRegisterUser
		mockRepository func(*mocks.UserRepository, *mocks.Bcrypt, dtos.RequestRegisterUser)
		want           *dtos.ResponseDataUser
		err            error
		wantErr        bool
	}{
		{
			name: "should successfully register a new user with service user",
			reqBody: dtos.RequestRegisterUser{
				Name:     nameUser,
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, reqBody dtos.RequestRegisterUser) {
				mockUser.On("IsEmailAlreadyRegistered", context.Background(), reqBody.Email).Return(false)
				mockBcrypt.On("HashPassword", reqBody.Password, 10).Return([]byte(hashPwd), nil)
				mockUser.On("PostUser", context.Background(), reqBody, hashPwd).Return(resUser, nil)
			},
			want: &dtos.ResponseDataUser{
				Name:      &nameUser,
				Email:     &emailUser,
				CreatedAt: &time.Time{},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "should error register a new user when user input email already exists",
			reqBody: dtos.RequestRegisterUser{
				Name:     nameUser,
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, reqBody dtos.RequestRegisterUser) {
				mockUser.On("IsEmailAlreadyRegistered", context.Background(), reqBody.Email).Return(true)
			},
			want:    nil,
			err:     apperrors.ErrUserEmailAlreadyExists,
			wantErr: true,
		},
		{
			name: "should error register a new user when something wrong in hashing password",
			reqBody: dtos.RequestRegisterUser{
				Name:     nameUser,
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, reqBody dtos.RequestRegisterUser) {
				mockUser.On("IsEmailAlreadyRegistered", context.Background(), reqBody.Email).Return(false)
				mockBcrypt.On("HashPassword", reqBody.Password, 10).Return(nil, apperrors.ErrUserFailedRegister)
			},
			want:    nil,
			err:     apperrors.ErrUserFailedRegister,
			wantErr: true,
		},
		{
			name: "should error register a new user when something wrong in post user repository",
			reqBody: dtos.RequestRegisterUser{
				Name:     nameUser,
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, reqBody dtos.RequestRegisterUser) {
				mockUser.On("IsEmailAlreadyRegistered", context.Background(), reqBody.Email).Return(false)
				mockBcrypt.On("HashPassword", reqBody.Password, 10).Return([]byte(hashPwd), nil)
				mockUser.On("PostUser", context.Background(), reqBody, hashPwd).Return(nil, apperrors.ErrUserFailedRegister)
			},
			want:    nil,
			err:     apperrors.ErrUserFailedRegister,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUser := &mocks.UserRepository{}
			mockBcrypt := &mocks.Bcrypt{}
			mockJwt := &mocks.JWTProvider{}
			userService := services.NewUserServiceImplementation(mockUser, mockBcrypt, mockJwt)
			tt.mockRepository(mockUser, mockBcrypt, tt.reqBody)

			user, err := userService.PostRegisterUserService(context.Background(), tt.reqBody)

			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.want.Name, user.Name)
				assert.Equal(t, tt.want.Email, user.Email)
			}
		})
	}
}

func TestUserServiceImplementation_PostLoginUserService(t *testing.T) {
	var (
		nameUser     = "Frierennnnnnn"
		emailUser    = "frieren@gmail.com"
		passwordUser = "12345"
		accesToken   = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsaWJyYXJ5X2FwaV91c2VyIiwiZXhwIjoxNzIwMTc2MzUwLCJpYXQiOjE3MjAxNzI3NTAsInVpZCI6NH0.V7BpiDtlQnIsn5Iu9p_LXTeVMIg5CGevtdu33F6kGRY"
	)
	resUser := &models.User{
		ID:        1,
		Name:      nameUser,
		Email:     emailUser,
		Password:  passwordUser,
		CreatedAt: time.Time{},
		UpdatedAt: &time.Time{},
		DeleteAt:  nil,
	}
	tests := []struct {
		name           string
		reqBody        dtos.RequestLoginUser
		mockRepository func(*mocks.UserRepository, *mocks.Bcrypt, *mocks.JWTProvider, dtos.RequestLoginUser)
		want           *dtos.ResponseDataUser
		err            error
		wantErr        bool
	}{
		{
			name: "should successfully login a user with service user",
			reqBody: dtos.RequestLoginUser{
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, mockJwt *mocks.JWTProvider, reqBody dtos.RequestLoginUser) {
				mockUser.On("GetUserByEmail", context.Background(), reqBody.Email).Return(resUser, nil)
				mockBcrypt.On("CheckPassword", reqBody.Password, []byte(resUser.Password)).Return(true, nil)
				mockJwt.On("CreateToken", int64(resUser.ID)).Return(accesToken, nil)
			},
			want: &dtos.ResponseDataUser{
				AccessToken: &accesToken,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "should error login a user when email system does not recognize the email",
			reqBody: dtos.RequestLoginUser{
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, mockJwt *mocks.JWTProvider, reqBody dtos.RequestLoginUser) {
				mockUser.On("GetUserByEmail", context.Background(), reqBody.Email).Return(nil, apperrors.ErrUserInvalidEmailPassword)
			},
			want:    nil,
			err:     apperrors.ErrUserInvalidEmailPassword,
			wantErr: true,
		},
		{
			name: "should error login a user when invalid password",
			reqBody: dtos.RequestLoginUser{
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, mockJwt *mocks.JWTProvider, reqBody dtos.RequestLoginUser) {
				mockUser.On("GetUserByEmail", context.Background(), reqBody.Email).Return(resUser, nil)
				mockBcrypt.On("CheckPassword", reqBody.Password, []byte(resUser.Password)).Return(false, apperrors.ErrUserInvalidEmailPassword)
			},
			want:    nil,
			err:     apperrors.ErrUserInvalidEmailPassword,
			wantErr: true,
		},
		{
			name: "should error login a user when something wrong when system fail make a new token",
			reqBody: dtos.RequestLoginUser{
				Email:    emailUser,
				Password: passwordUser,
			},
			mockRepository: func(mockUser *mocks.UserRepository, mockBcrypt *mocks.Bcrypt, mockJwt *mocks.JWTProvider, reqBody dtos.RequestLoginUser) {
				mockUser.On("GetUserByEmail", context.Background(), reqBody.Email).Return(resUser, nil)
				mockBcrypt.On("CheckPassword", reqBody.Password, []byte(resUser.Password)).Return(true, nil)
				mockJwt.On("CreateToken", int64(resUser.ID)).Return("", apperrors.ErrISE)
			},
			want:    nil,
			err:     apperrors.ErrISE,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUser := &mocks.UserRepository{}
			mockBcrypt := &mocks.Bcrypt{}
			mockJwt := &mocks.JWTProvider{}
			userService := services.NewUserServiceImplementation(mockUser, mockBcrypt, mockJwt)
			tt.mockRepository(mockUser, mockBcrypt, mockJwt, tt.reqBody)

			generateToken, err := userService.PostLoginUserService(context.Background(), tt.reqBody)

			if tt.wantErr {
				assert.Equal(t, tt.err, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, generateToken)
				assert.Equal(t, *tt.want.AccessToken, *generateToken.AccessToken)
			}
		})
	}
}
