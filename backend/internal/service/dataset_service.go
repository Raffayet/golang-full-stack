package service

import (
	"context"
	"data-merging/internal/domain"
	"data-merging/internal/repository"
)

// DatasetService defines methods for working with profiles
type DatasetService interface {
	GetDatasets(ctx context.Context) ([]domain.Dataset, error)
}

// DatasetServiceImpl implements the ProfileService interface
type DatasetServiceImpl struct {
	repo repository.DatasetRepository
}

func NewDatasetService(repo repository.DatasetRepository) *DatasetServiceImpl {
	return &DatasetServiceImpl{repo: repo}
}

func (s *DatasetServiceImpl) GetDatasets(ctx context.Context) ([]domain.Dataset, error) {
	return s.repo.FetchDatasets(ctx)
}
