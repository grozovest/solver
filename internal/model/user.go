package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type User struct {
	ID         uuid.UUID       `json:"id" db:"id"`
	FirstName  string          `json:"first_name" db:"first_name"`
	SecondName string          `json:"second_name" db:"second_name"`
	FatherName string          `json:"father_name" db:"father_name"`
	GroupName  string          `json:"group_name" db:"group_name"`
	Password   string          `json:"password" db:"password"`
	Balance    decimal.Decimal `json:"balance" db:"balance"`
}
