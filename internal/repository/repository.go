package repository

import (
	"analytics-service/internal/domain"
	"context"
)

type ClickRepository interface {
	SaveClick(ctx context.Context, event domain.ClickEvent) error
}