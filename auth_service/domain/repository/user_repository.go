package repository

import (
	"auth-service/errors"
	"auth-service/models"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, *errors.AppError)
}
