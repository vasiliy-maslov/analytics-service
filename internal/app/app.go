package app

import (
	"analytics-service/internal/config"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nats-io/nats.go"
)

type App struct {
	db   *sql.DB
	nats *nats.Conn
	log  *slog.Logger
}

func New(log *slog.Logger, cfg *config.Config) (*App, error) {
	db, err := sql.Open("pgx", cfg.Postgres.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	log.Info("Successfully connected to PostgreSQL")

	natsConn, err := nats.Connect(cfg.NATS.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %w", err)
	}
	log.Info("Successfully connected to NATS")

	return &App{
		db:   db,
		nats: natsConn,
		log:  log,
	}, nil
}

func (a *App) Run() {
	fmt.Println("App starting...")
}

func (a *App) Stop() {
	a.log.Info("App stopped...")
	a.db.Close()
	a.nats.Close()
}
