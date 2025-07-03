package services

import (
	"fmt"
	"user-service/models"
	"user-service/repository"
	"user-service/utils"
)

type UserService struct {
	UserRepo *repository.UserRepo
}

func (s *UserService) RegisterUser(email, password string) (int64, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return 0, err
	}

	user := models.User{Email: email, Password: hashed}
	
	userId, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *UserService) LoginUser(email, password string) (string, error) {
	user, err := s.UserRepo.GetUser(email)

	if err != nil {
		return "", err
	}

	if !utils.VerfyPassword(password, user.Password) {
		return "", fmt.Errorf("password verfication failed")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}