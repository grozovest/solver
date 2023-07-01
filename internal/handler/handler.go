package handler

import (
	"context"

	"solver/internal/model"
	"solver/pkg/logger"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	logger  *logger.Logger
	useCase useCase
}

type useCase interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error)
	IncreaseUserBalance(ctx context.Context, userID uuid.UUID, sum decimal.Decimal) (*model.User, error)
	SolveTask(ctx context.Context, userID uuid.UUID, task []uint32) ([]uint32, error)
}

func New(logger *logger.Logger, useCase useCase) *Handler {
	return &Handler{
		logger:  logger,
		useCase: useCase,
	}
}
