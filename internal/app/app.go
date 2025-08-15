package app

import (
	"analytics-service/internal/config"
	"analytics-service/internal/repository/postgres"
	"analytics-service/internal/service"
	natsConsumer "analytics-service/internal/transport/nats"
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nats-io/nats.go"
)

type App struct {
	db       *sql.DB
	nats     *nats.Conn
	log      *slog.Logger
	consumer *natsConsumer.Consumer
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

	clickRepo := postgres.NewClickRepository(db)
	analyticsService := service.NewAnalyticsService(clickRepo)

	consumer, err := natsConsumer.NewConsumer(natsConn, analyticsService, log.With(slog.String("component", "nats_consimer")))
	if err != nil {
		return nil, fmt.Errorf("failed to create nats consumer: %w", err)
	}


	return &App{
		db:   db,
		nats: natsConn,
		log:  log,
		consumer: consumer,
	}, nil
}

func (a *App) Run() {
	a.log.Info("Application starting...")
	if err := a.consumer.Start(); err != nil {
		a.log.Error("failed to start nats consumer", slog.Any("error", err))
	}
}

func (a *App) Stop() {
	a.log.Info("App stopped...")
	a.db.Close()
	a.nats.Close()
}
