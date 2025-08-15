package postgres

import (
	"analytics-service/internal/domain"
	"context"
	"database/sql"
	"fmt"
)

type ClickRepository struct {
	db *sql.DB
}

func NewClickRepository(db *sql.DB) *ClickRepository {
	return &ClickRepository{db: db}
}

func (r *ClickRepository) SaveClick(ctx context.Context, event domain.ClickEvent) error {
	query := `INSERT INTO clicks (alias, timestamp, source) VALUES ($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, event.Alias, event.Timestamp, event.Source)
	if err != nil {
		return fmt.Errorf("failed to save click to postgres: %w", err)
	}

	return nil
}