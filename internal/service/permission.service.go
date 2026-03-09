package service

import (
	dto "base-api/internal/dto/permission"
	"base-api/internal/model"
	"base-api/internal/repository"
	"context"

	"github.com/google/uuid"
)

type PermissionService interface {
	GetAll(ctx context.Context) ([]dto.GroupedPermissionResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.PermissionResponse, error)
	Create(ctx context.Context, req dto.PermissionCreateRequest) error
	Update(ctx context.Context, id uuid.UUID, req dto.PermissionUpdateRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type permissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &permissionService{repo}
}

func (s *permissionService) GetAll(ctx context.Context) ([]dto.GroupedPermissionResponse, error) {
	permissions, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	groupedMap := make(map[string][]dto.PermissionResponse)
	var groupOrder []string

	for _, p := range permissions {
		groupName := "Others"
		if p.Group != nil && *p.Group != "" {
			groupName = *p.Group
		}

		if _, exists := groupedMap[groupName]; !exists {
			groupOrder = append(groupOrder, groupName)
		}

		groupedMap[groupName] = append(groupedMap[groupName], dto.PermissionResponse{
			ID:        p.ID,
			Name:      p.Name,
			Code:      p.Code,
			Group:     p.Group,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	var res []dto.GroupedPermissionResponse
	for _, groupName := range groupOrder {
		res = append(res, dto.GroupedPermissionResponse{
			Group: groupName,
			Items: groupedMap[groupName],
		})
	}

	return res, nil
}

func (s *permissionService) GetByID(ctx context.Context, id uuid.UUID) (dto.PermissionResponse, error) {
	permission, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.PermissionResponse{}, err
	}

	return dto.PermissionResponse{
		ID:        permission.ID,
		Name:      permission.Name,
		Code:      permission.Code,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}, nil
}
func (s *permissionService) Create(ctx context.Context, req dto.PermissionCreateRequest) error {
	p := model.Permission{
		Name:  req.Name,
		Code:  req.Code,
		Group: &req.Group,
	}
	_, err := s.repo.UpsertPermission(ctx, nil, p)
	return err
}

func (s *permissionService) Update(ctx context.Context, id uuid.UUID, req dto.PermissionUpdateRequest) error {
	p, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	p.Name = req.Name
	p.Code = req.Code
	p.Group = &req.Group
	_, err = s.repo.UpsertPermission(ctx, nil, p)
	return err
}

func (s *permissionService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
