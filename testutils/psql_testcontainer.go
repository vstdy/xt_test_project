package testutils

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// PostgreSQLContainer wraps testcontainers.Container with extra methods.
type (
	PostgreSQLContainer struct {
		testcontainers.Container
		Config PostgreSQLContainerConfig
	}

	PostgreSQLContainerOption func(c *PostgreSQLContainerConfig)

	PostgreSQLContainerConfig struct {
		ImageTag   string
		User       string
		Password   string
		MappedPort string
		Database   string
	}
)

// GetDSN returns DB connection URL.
func (c PostgreSQLContainer) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", c.Config.User, c.Config.Password, c.Config.MappedPort, c.Config.Database)
}

func WithPostgreSQLTag(tag string) PostgreSQLContainerOption {
	return func(c *PostgreSQLContainerConfig) {
		c.ImageTag = tag
	}
}

func WithPostgreSQLDatabaseName(dbName string) PostgreSQLContainerOption {
	return func(c *PostgreSQLContainerConfig) {
		c.Database = dbName
	}
}

// NewPostgreSQLContainer creates and starts a PostgreSQL container.
func NewPostgreSQLContainer(ctx context.Context, opts ...PostgreSQLContainerOption) (*PostgreSQLContainer, error) {
	const (
		psqlImage = "postgres"
		psqlPort  = "5432"
	)

	config := PostgreSQLContainerConfig{
		ImageTag: "12.10-alpine",
		User:     "user",
		Password: "password",
		Database: "mockdb",
	}
	for _, opt := range opts {
		opt(&config)
	}

	containerPort := psqlPort + "/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Env: map[string]string{
				"POSTGRES_USER":     config.User,
				"POSTGRES_PASSWORD": config.Password,
				"POSTGRES_DB":       config.Database,
			},
			ExposedPorts: []string{
				containerPort,
			},
			Image:      fmt.Sprintf("%s:%s", psqlImage, config.ImageTag),
			SkipReaper: true,
			WaitingFor: wait.ForSQL(
				nat.Port(containerPort),
				"postgres",
				func(port nat.Port) string {
					return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", config.User, config.Password, port.Port(), config.Database)
				}).
				Timeout(5 * time.Second),
		},
		ProviderType: testcontainers.ProviderDocker,
	}

	provider, err := req.ProviderType.GetProvider()
	if err != nil {
		return nil, fmt.Errorf("getting request provider: %w", err)
	}

	container, err := provider.CreateContainer(ctx, req.ContainerRequest)
	if err != nil {
		return nil, fmt.Errorf("creating container: %w", err)
	}

	termCont := func() {
		termCtx, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer ctxCancel()
		container.Terminate(termCtx)
	}

	if err = container.Start(ctx); err != nil {
		termCont()
		return nil, fmt.Errorf("starting container: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, nat.Port(containerPort))
	if err != nil {
		termCont()
		return nil, fmt.Errorf("getting mapped port for (%s): %w", containerPort, err)
	}
	config.MappedPort = mappedPort.Port()

	return &PostgreSQLContainer{
		Container: container,
		Config:    config,
	}, nil
}
