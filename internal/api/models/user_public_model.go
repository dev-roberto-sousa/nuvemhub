package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UserPublic struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username" validate:"required"`
	Email    string    `json:"email" validate:"required"`
}

func (u *UserPublic) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
