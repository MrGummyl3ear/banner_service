package repository

import (
	"banner_service/pkg/model"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User, role string) (int, error) {
	var err error
	var userRole model.Role

	tx := r.db.Begin()

	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	userRole.Username = user.Username
	userRole.Role = role
	err = tx.Create(&userRole).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return user.Id, tx.Commit().Error
}

func (r *AuthPostgres) GetUser(username, password string) ([]string, error) {
	var user model.User
	var roles []model.Role
	tx := r.db.Begin()

	req := tx.Where("username = ? and password = ?", username, password).First(&user)
	if req.Error != nil {
		tx.Rollback()
		return nil, req.Error
	}

	req = tx.Where("username = ?", username).Find(&roles)
	if req.Error != nil {
		tx.Rollback()
		return nil, req.Error
	}
	var list []string
	for _, user := range roles {
		list = append(list, user.Role)
	}
	return list, tx.Commit().Error
}
