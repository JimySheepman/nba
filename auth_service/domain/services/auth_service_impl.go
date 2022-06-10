package services

import (
	"auth-service/domain/repository"
	"auth-service/errors"
	"auth-service/models"
)

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func NewAuthServiceImpl(userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{userRepository: userRepository}
}

func (asi AuthServiceImpl) GetUserByEmail(email string) (*models.User, *errors.AppError) {
	result, err := asi.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return result, nil
}
