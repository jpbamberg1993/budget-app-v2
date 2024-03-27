package utils

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"jemyishai/budget-app/internal/datasources/data"
	"os"
)

func SetupPostgres() (*data.Queries, error) {
	ctx := context.Background()
	fmt.Printf("POSTGRES_USER: %s\n", os.Getenv("POSTGRES_USER"))
	fmt.Printf("POSTGRES_PASSWORD: %s\n", os.Getenv("POSTGRES_PASSWORD"))
	conn, err := pgx.Connect(
		ctx,
		fmt.Sprintf("postgres://%s:%s@localhost:5432/budgetdb",
			"budgetuser",
			"budgetpassword"))
	if err != nil {
		zap.L().Fatal("failed to connect to postgres", zap.Error(err))
	}
	//defer conn.Close(ctx)
	queries := data.New(conn)
	return queries, nil
}
