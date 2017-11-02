package services

import "context"

type GenericService interface {
	GetGeneric(ctx context.Context, filters *GenericFilters) ([]*Generic, error)
}

type genericService struct {
	repository GenericRepository
}

func NewGenericService(repository GenericRepository) GenericService {
	return &genericService{
		repository: repository,
	}
}

func (p *genericService) GetGeneric(ctx context.Context, filters *GenericFilters) ([]*Generic, error) {
	return p.repository.GetGeneric(ctx, filters)
}
