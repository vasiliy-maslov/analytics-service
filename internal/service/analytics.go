package service

import (
	"analytics-service/internal/domain"
	"analytics-service/internal/repository"
	"context"
)

type AnalyticsService interface {
	HandleClickEvent(ctx context.Context, event domain.ClickEvent) error
}

type service struct {
	repo repository.ClickRepository
}

func NewAnalyticsService(repo repository.ClickRepository) *service {
	return &service{repo: repo}
}

func (s *service) HandleClickEvent(ctx context.Context, event domain.ClickEvent) error {
	return s.repo.SaveClick(ctx, event)
}