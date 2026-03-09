package service

import (
	dtoRole "base-api/internal/dto/role"
	dto "base-api/internal/dto/user"
	"base-api/internal/model"
	"base-api/internal/pkg/helper"
	"base-api/internal/repository"
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserService interface {
	GetAll(ctx context.Context, page, limit int) ([]dto.UserResponse, int, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.UserResponse, error)
	Create(ctx context.Context, req dto.UserCreateRequest) error
	Update(ctx context.Context, id uuid.UUID, req dto.UserUpdateRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
	db   *sqlx.DB
}

func NewUserService(repo repository.UserRepository, db *sqlx.DB) UserService {
	return &userService{repo, db}
}

func (s *userService) GetAll(ctx context.Context, page, limit int) ([]dto.UserResponse, int, error) {
	offset := (page - 1) * limit

	users, err := s.repo.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	res := make([]dto.UserResponse, 0, len(users))

	for _, user := range users {
		var roles []dtoRole.RoleResponse
		if err := json.Unmarshal(user.Roles, &roles); err != nil {
			return nil, 0, err
		}

		res = append(res, dto.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Roles: roles,
		})
	}

	return res, total, nil
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (dto.UserResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	res := dto.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}

func (s *userService) Create(ctx context.Context, req dto.UserCreateRequest) error {
	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := model.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: hashedPassword,
	}

	_, err = s.repo.UpsertUser(ctx, nil, user)
	return err
}

func (s *userService) Update(ctx context.Context, id uuid.UUID, req dto.UserUpdateRequest) error {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Name != nil {
		user.Name = req.Name
	}
	if req.Password != nil {
		hashedPassword, err := helper.HashPassword(*req.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	_, err = s.repo.UpsertUser(ctx, nil, user)
	return err
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
