package service

import (
	dto "base-api/internal/dto/auth"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/repository"
	"context"
	"errors"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := s.repo.FindAuthUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return "", constants.ErrUserNotFound
		}
		return "", err
	}

	if !helper.CheckPassword(user.Password, req.Password) {
		return "", constants.ErrInvalidCredential
	}

	token, err := helper.GenerateToken(user.ID, req.Email, user.Roles, user.Permissions)
	if err != nil {
		return "", err
	}

	return token, nil
}
