# Analytics Service

A microservice responsible for consuming and storing click analytics data. This project is a practical implementation of Clean Architecture, built in Go.

This service acts as a subscriber to a NATS message broker, listening for click events published by other systems (like the `url-shortener` service). Upon receiving an event, it persists the data into its own isolated PostgreSQL database for future analysis.

This project is the second part of a comprehensive Go mentorship program, focusing on building interconnected microservices and deepening understanding of backend principles.

## Technology Stack

- **Language:** Go
- **Architecture:** Clean Architecture
- **Database:** PostgreSQL
- **Message Broker:** NATS (JetStream Subscriber)
- **Primary Role:** Asynchronous Event Consumer
- **Production-Ready Features:** Structured Configuration (`cleanenv`), Graceful Shutdown, Structured Logging (`slog`), Database Migrations.

## Architecture

The service is structured to be highly maintainable and testable:
- **`cmd/app`:** The entrypoint of the application.
- **`internal/config`:** Handles structured configuration loading from YAML files and environment variables.
- **`internal/app`:** The composition root, responsible for initializing and wiring all components (DI).
- **`internal/domain`:** Contains the core business entities (e.g., `ClickEvent`).
- **`internal/service`:** Implements the core business logic (Use Cases).
- **`internal/repository`:** Defines interfaces for data storage and provides PostgreSQL implementation.
- **`internal/transport/nats`:** The inbound adapter responsible for consuming messages from the NATS broker.

## How to Run

> **Prerequisites:**
> 1. A running PostgreSQL database.
> 2. A running NATS JetStream server.
> 3. The `url-shortener` service (or any other publisher) must have created the `CLICKS` stream on NATS.

The easiest way to run the required infrastructure is by using the `docker-compose.yml` from the [url-shortener-service](https://github.com/vasiliy-maslov/url-shortener-service) project.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/analytics-service.git
    cd analytics-service
    ```

2.  **Configure the application:**
    Ensure your `config/config.yaml` points to the correct NATS and PostgreSQL instances. The defaults are set to work with the `url-shortener`'s Docker Compose setup.

3.  **Run the service:**
    ```bash
    go run ./cmd/app/main.go
    ```
    The service will start, connect to the database and NATS, and begin listening for click events.