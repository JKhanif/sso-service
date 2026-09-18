package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/JKhanif/sso-service/internal/app/grpc"
	authsvc "github.com/JKhanif/sso-service/internal/service/auth"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: инициализировать хранилище.

	// TODO: инициализировать auth-сервис, когда появится хранилище.

	authService := authsvc.New(log)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}