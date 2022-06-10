package services

import (
	"auth-service/errors"
	"auth-service/models"
)

type AuthService interface {
	GetUserByEmail(email string) (*models.User, *errors.AppError)
}
