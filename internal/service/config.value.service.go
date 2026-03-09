package service

import (
	dto "base-api/internal/dto/config"
	"base-api/internal/model"
	"base-api/internal/repository"
	"context"

	"github.com/google/uuid"
)

type ConfigValueService interface {
	GetValuesByConfigID(ctx context.Context, configID uuid.UUID) ([]dto.ConfigValueResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.ConfigValueResponse, error)
	Create(ctx context.Context, configID uuid.UUID, req dto.ConfigValueCreateRequest) error
	Update(ctx context.Context, id uuid.UUID, req dto.ConfigValueUpdateRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type configValueService struct {
	repo repository.ConfigValueRepository
}

func NewConfigValueService(repo repository.ConfigValueRepository) ConfigValueService {
	return &configValueService{repo}
}

func (s *configValueService) GetValuesByConfigID(ctx context.Context, configID uuid.UUID) ([]dto.ConfigValueResponse, error) {
	values, err := s.repo.FindValuesByConfigID(ctx, configID)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ConfigValueResponse, 0, len(values))
	for _, val := range values {
		res = append(res, dto.ConfigValueResponse{
			ID:        val.ID,
			ConfigID:  val.ConfigID,
			Name:      val.Name,
			Code:      val.Code,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return res, nil
}

func (s *configValueService) GetByID(ctx context.Context, id uuid.UUID) (dto.ConfigValueResponse, error) {
	val, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return dto.ConfigValueResponse{}, err
	}

	return dto.ConfigValueResponse{
		ID:        val.ID,
		ConfigID:  val.ConfigID,
		Name:      val.Name,
		Code:      val.Code,
		CreatedAt: val.CreatedAt,
		UpdatedAt: val.UpdatedAt,
	}, nil
}
func (s *configValueService) Create(ctx context.Context, configID uuid.UUID, req dto.ConfigValueCreateRequest) error {
	val := model.ConfigValues{
		ConfigID: configID,
		Name:     req.Name,
		Code:     req.Code,
	}
	return s.repo.UpsertConfigValue(ctx, nil, val)
}

func (s *configValueService) Update(ctx context.Context, id uuid.UUID, req dto.ConfigValueUpdateRequest) error {
	val, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	val.Name = req.Name
	val.Code = req.Code
	return s.repo.UpsertConfigValue(ctx, nil, val)
}

func (s *configValueService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
