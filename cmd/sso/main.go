package main

import (
	"fmt"

	"github.com/JKhanif/sso-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("%+v\n", cfg)
	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}
