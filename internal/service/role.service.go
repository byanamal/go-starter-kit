package service

import (
	dto "base-api/internal/dto/role"
	"base-api/internal/model"
	"base-api/internal/repository"
	"context"

	"github.com/google/uuid"
)

type RoleService interface {
	GetAll(ctx context.Context) ([]dto.RoleResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.RoleResponse, error)
	Create(ctx context.Context, req dto.RoleCreateRequest) error
	Update(ctx context.Context, id uuid.UUID, req dto.RoleUpdateRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo}
}

func (s *roleService) GetAll(ctx context.Context) ([]dto.RoleResponse, error) {
	roles, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var res []dto.RoleResponse
	for _, role := range roles {
		res = append(res, dto.RoleResponse{
			ID:          role.ID,
			Name:        role.Name,
			Code:        role.Code,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
			Permissions: role.Permissions,
		})
	}

	return res, nil
}

func (s *roleService) GetByID(ctx context.Context, id uuid.UUID) (dto.RoleResponse, error) {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.RoleResponse{}, err
	}

	return dto.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
			Permissions: role.Permissions,
	}, nil
}
func (s *roleService) Create(ctx context.Context, req dto.RoleCreateRequest) error {
	role := model.Role{
		Name: req.Name,
		Code: req.Code,
	}
	_, err := s.repo.UpsertRole(ctx, nil, role)
	return err
}

func (s *roleService) Update(ctx context.Context, id uuid.UUID, req dto.RoleUpdateRequest) error {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	role.Name = req.Name
	role.Code = req.Code
	_, err = s.repo.UpsertRole(ctx, nil, role)
	return err
}

func (s *roleService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
