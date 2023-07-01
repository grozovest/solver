package usecase

import (
	"context"

	"solver/internal/model"
	"solver/pkg/logger"

	"github.com/google/uuid"
)

type storage interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type service struct {
	logger  *logger.Logger
	storage storage
}

func New(logger *logger.Logger, storage storage) *service {
	return &service{
		logger:  logger,
		storage: storage,
	}
}
