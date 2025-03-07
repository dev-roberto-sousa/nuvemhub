package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleType string

const (
	Student RoleType = "student"
	Teacher RoleType = "teacher"
	Admin   RoleType = "admin"
)

type User struct {
	ID        uuid.UUID      `json:"uuid" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string         `json:"name" gorm:"size:100;not null"`
	Username  string         `json:"username" gorm:"size:100;unique;not null"`
	Email     string         `json:"email" gorm:"size:100;unique;not null"`
	Password  string         `json:"password" gorm:"size:255;not null"`
	Role      RoleType       `json:"role" gorm:"size:20;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
