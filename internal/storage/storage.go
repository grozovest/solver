package storage

import (
	"solver/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	db *pgxpool.Pool
	lg *logger.Logger
}

func New(db *pgxpool.Pool, lg *logger.Logger) *storage {
	return &storage{
		db: db,
		lg: lg,
	}
}
