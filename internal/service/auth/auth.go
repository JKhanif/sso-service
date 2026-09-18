package auth

import (
	"context"
	"errors"
	"log/slog"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
)

type Service struct {
	log *slog.Logger
	// TODO: хранилище (storage)
}

func New(log *slog.Logger) *Service {
	return &Service{log: log}
}

func (s *Service) Login(ctx context.Context, email string, password string, appID int) (token string, err error) {
	// TODO: реализовать.
	return "", ErrInvalidCredentials
}

func (s *Service) RegisterNewUser(ctx context.Context, email string, password string) (userID int64, err error) {
	// TODO: реализовать.
	return 0, errors.New("not implemented")
}

func (s *Service) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	// TODO: реализовать.
	return false, errors.New("not implemented")
}