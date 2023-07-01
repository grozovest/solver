package storage

import (
	"context"
	"fmt"
	"time"

	"solver/internal/model"

	"github.com/google/uuid"
)

const defaultTimeout = time.Millisecond * 500

func (s *storage) CreateUser(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `INSERT INTO users (first_name, second_name, father_name, group_name, balance, password) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(ctx, query, user.FirstName, user.SecondName, user.FatherName, user.GroupName, user.Balance, user.Password)
	if err != nil {
		return err
	}

	return err
}

func (s *storage) GetUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	var user *model.User

	query := fmt.Sprintf(`SELECT first_name, second_name, father_name, group_name, balance, password FROM users WHERE id = '%s'`, userID)

	err := s.db.QueryRow(ctx, query).Scan(&user.FirstName, &user.SecondName, &user.FatherName, &user.GroupName, &user.Balance, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *storage) UpdateUser(ctx context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := fmt.Sprintf(`UPDATE users SET balance = $1 WHERE id = '%s'`, user.Balance, user.ID)

	err := s.db.QueryRow(ctx, query).Scan(&user.Balance)
	if err != nil {
		return err
	}

	return err
}
