package repositories

import (
	"github.com/dev-roberto-sousa/nuvemhub/internal/api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
