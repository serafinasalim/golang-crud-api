package service

import (
	"golang-crud-api/dto"
	"golang-crud-api/helper"
	"golang-crud-api/model"
	"golang-crud-api/repository"
	"time"
)

type AuthService struct {
	repository *repository.AuthRepository
}

func NewAuthService(repository *repository.AuthRepository) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Register(params dto.RegisterRequest) (err error) {
	// Hash the password
	hashedPassword, err := helper.HashPassword(params.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Username:  params.Username,
		Email:     params.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	}

	err = s.repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(params dto.LoginRequest) (result *dto.LoginResponse, err error) {
	// Find user by email
	user, err := s.repository.GetUserByUsername(params.Username)
	if err != nil {
		return nil, err
	}

	// Verify password
	err = helper.VerifyPassword(user.Password, params.Password)
	if err != nil {
		return nil, err
	}

	// Generate JWT token
	token, err := helper.GenerateJWT(user.Id)
	if err != nil {
		return nil, err
	}

	// Return response with token
	result = &dto.LoginResponse{
		Username: user.Username,
		Token:    token,
	}

	return result, nil
}
