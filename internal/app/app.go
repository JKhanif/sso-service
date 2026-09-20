package app

import (
	"context"
	"log/slog"
	"time"

	grpcapp "github.com/JKhanif/sso-service/internal/app/grpc"
	authsvc "github.com/JKhanif/sso-service/internal/services/auth"
	"github.com/JKhanif/sso-service/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, databaseURL string, tokenTTL time.Duration) *App {
	storage, err := postgres.New(context.Background(), databaseURL)
	if err != nil {
		panic(err)
	}

	authService := authsvc.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}