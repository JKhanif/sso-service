package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

// Конфигурация загружается из config.yaml, структура повторяет его поля
type Config struct {
	Env         string        `yaml:"env" env:"ENV" env-default:"local"`
	DatabaseURL string        `yaml:"database_url" env:"DATABASE_URL" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env:"TOKEN_TTL" env-default:"3600s"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env:"GRPC_PORT" env-default:"50051"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"10s"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	return MustLoadByPath(path)
}

func MustLoadByPath(configPath string) *Config {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		panic("config file does not found exist: " + configPath)
	}

	var cfg Config

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	// Путь к конфигу задаётся одним из двух способов:
	// 1. Через флаг командной строки: sso --config="path/to/config.yaml"
	// 2. Через переменную окружения CONFIG_PATH - она задаётся в системе/терминале перед запуском
	// Приоритет у флага: если задан --config, он перекрывает CONFIG_PATH
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	// Подгружаем .env, чтобы можно было положить CONFIG_PATH туда
	_ = godotenv.Load()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
