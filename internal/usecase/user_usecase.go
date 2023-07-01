package usecase

import (
	"context"
	"errors"

	"solver/internal/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	minUserBalance = -1000
	solvingCost    = 200
)

func (s *service) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := validateUser(user); err != nil {
		return nil, err
	}

	user.Password = getPasswordHash(user.Password)
	user.ID = uuid.New()

	err := s.storage.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.storage.GetUser(ctx, user.ID)
}

func (s *service) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return s.storage.GetUser(ctx, userID)
}

func (s *service) IncreaseUserBalance(ctx context.Context, userID uuid.UUID, sum decimal.Decimal) (*model.User, error) {
	user, err := s.storage.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	user.Balance.Add(sum)

	err = s.storage.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.storage.GetUser(ctx, user.ID)
}

func (s *service) SolveTask(ctx context.Context, userId uuid.UUID, task []uint32) ([]uint32, error) {
	user, err := s.storage.GetUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	if user.Balance.LessThan(decimal.New(minUserBalance, 1)) {
		return nil, errors.New("not enough money on user's balance")
	}

	result, err := findMissingNumbers(task)
	if err != nil {
		return nil, err
	}

	user.Balance.Sub(decimal.New(solvingCost, 1))

	err = s.storage.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
