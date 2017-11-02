package controllers

import (
	"context"
	"net/http"

	"github.com/rewardStyle/generic-service/services"
)

type GenericController struct {
	service services.GenericService
}

func NewGenericController(svc services.GenericService) *GenericController {
	return &GenericController{
		service: svc,
	}
}

func (p *GenericController) GetGeneric(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	generics, err := p.service.GetGeneric(ctx, &services.GenericFilters{})
	if err != nil {
		RenderError(ctx, rw, err)
		return
	}

	Render(ctx, rw, 200, generics)
}
