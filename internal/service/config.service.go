package service

import (
	dto "base-api/internal/dto/config"
	"base-api/internal/model"
	"base-api/internal/repository"
	"context"

	"github.com/google/uuid"
)

type ConfigService interface {
	// Config Methods
	GetAll(ctx context.Context) ([]dto.ConfigResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (dto.ConfigResponse, error)
	Create(ctx context.Context, req dto.ConfigCreateRequest) error
	Update(ctx context.Context, id uuid.UUID, req dto.ConfigUpdateRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type configService struct {
	repo repository.ConfigRepository
}

func NewConfigService(repo repository.ConfigRepository) ConfigService {
	return &configService{repo}
}

// Config Implementation

func (h *configService) GetAll(ctx context.Context) ([]dto.ConfigResponse, error) {
	configs, err := h.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ConfigResponse, 0, len(configs))
	for _, cfg := range configs {
		res = append(res, dto.ConfigResponse{
			ID:        cfg.ID,
			Name:      cfg.Name,
			Code:      cfg.Code,
			CreatedAt: cfg.CreatedAt,
			UpdatedAt: cfg.UpdatedAt,
		})
	}
	return res, nil
}

func (h *configService) GetByID(ctx context.Context, id uuid.UUID) (dto.ConfigResponse, error) {
	cfg, err := h.repo.FindByID(ctx, id)
	if err != nil {
		return dto.ConfigResponse{}, err
	}

	return dto.ConfigResponse{
		ID:        cfg.ID,
		Name:      cfg.Name,
		Code:      cfg.Code,
		CreatedAt: cfg.CreatedAt,
		UpdatedAt: cfg.UpdatedAt,
	}, nil
}
func (h *configService) Create(ctx context.Context, req dto.ConfigCreateRequest) error {
	cfg := model.Config{
		Name: req.Name,
		Code: req.Code,
	}
	_, err := h.repo.UpsertConfig(ctx, nil, cfg)
	return err
}

func (h *configService) Update(ctx context.Context, id uuid.UUID, req dto.ConfigUpdateRequest) error {
	cfg, err := h.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	cfg.Name = req.Name
	cfg.Code = req.Code
	_, err = h.repo.UpsertConfig(ctx, nil, cfg)
	return err
}

func (h *configService) Delete(ctx context.Context, id uuid.UUID) error {
	return h.repo.Delete(ctx, id)
}
