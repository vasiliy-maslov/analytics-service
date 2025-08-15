package nats

import (
	"analytics-service/internal/domain"
	"analytics-service/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/nats-io/nats.go"
)

type Consumer struct {
	service service.AnalyticsService
	js      nats.JetStreamContext
	log     *slog.Logger
}

func NewConsumer(natsConn *nats.Conn, service service.AnalyticsService, log *slog.Logger) (*Consumer, error) {
	js, err := natsConn.JetStream()
	if err != nil {
		return nil, fmt.Errorf("failed to get nats jetstream context: %w", err)
	}

	return &Consumer{
		service: service,
		js:      js,
		log:     log,
	}, nil
}

func (c *Consumer) Start() error {
	_, err := c.js.Subscribe("clicks.events", func(msg *nats.Msg) {
		msg.Ack()

		c.log.Info("Received a message", slog.String("subject", msg.Subject))

		var event domain.ClickEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			c.log.Error("failed to unmarshal message", slog.Any("error", err))
			return 
		}
		err := c.service.HandleClickEvent(context.Background(), event)
		if err != nil {
			c.log.Error("failed to handle click", slog.Any("error", err))
			return 
		}
	}, nats.Durable("analytics-service-consumer"), nats.ManualAck())

	if err != nil {
		return fmt.Errorf("failed to subscribe to jetstream: %w", err)
	}

	c.log.Info("Subscribed to clicks.events subject")
	return nil
}
