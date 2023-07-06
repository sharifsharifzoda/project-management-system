package repository

import (
	"github.com/sharifsharifzoda/project-management-system/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *models.User) (int, error) {
	tx := r.db.Create(&user)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return user.ID, nil
}

func (r *AuthPostgres) GetUser(email string) (user models.User, err error) {
	tx := r.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	return user, nil
}

func (r *AuthPostgres) IsEmailUsed(email string) bool {
	var user models.User
	tx := r.db.Where("email = ?", email).Find(&user)
	if tx.Error != nil {
		return false
	}

	if user.Email == "" {
		return false
	}

	return true
}
