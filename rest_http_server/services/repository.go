package services

import "context"

type GenericRepository interface {
	GetGeneric(ctx context.Context, filters *GenericFilters) ([]*Generic, error)
}

type genericRepository struct {
}

func NewGenericRepository() GenericRepository {
	return &genericRepository{}
}

func (p *genericRepository) GetGeneric(ctx context.Context, filters *GenericFilters) ([]*Generic, error) {
	return []*Generic{
		&Generic{Name: "test"},
	}, nil
}
