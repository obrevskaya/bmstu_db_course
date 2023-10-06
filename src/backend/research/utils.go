package main

import (
	"context"
	"fmt"
	"github.com/pressly/goose"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func SetupTestDatabase() (testcontainers.Container, *gorm.DB, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	containerReq := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       "testdb",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_USER":     "postgres",
		},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: pwd + "/deployments/data",
				},
				Target: "/data",
			},
		},
	}

	dbContainer, err := testcontainers.GenericContainer(
		context.Background(),
		testcontainers.GenericContainerRequest{
			ContainerRequest: containerReq,
			Started:          true,
		})
	if err != nil {
		return nil, nil, err
	}

	host, _ := dbContainer.Host(context.Background())
	port, _ := dbContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf("host=%s port=%d user=postgres password=postgres dbname=testdb sslmode=disable", host, port.Int())
	pureDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("gorm open: %w", err)
	}

	sqlDB, err := pureDB.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("get db: %w", err)
	}

	if err = goose.Up(sqlDB, "./deployments/migrations/research_migrations"); err != nil {
		return nil, nil, fmt.Errorf("up migrations: %w", err)
	}

	//text, err := os.ReadFile("./deployments/res/10.sql")
	//if err != nil {
	//	return nil, nil, fmt.Errorf("read file: %w", err)
	//}
	//
	//if err := pureDB.Exec(string(text)).Error; err != nil {
	//	return nil, nil, fmt.Errorf("exec: %w", err)
	//}

	return dbContainer, pureDB, nil
}
