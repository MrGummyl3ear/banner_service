package repository

import (
	"banner_service/pkg/model"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user model.User, role string) (int, error)
	GetUser(username, password string) ([]string, error)
}

type Banner interface {
	CreateBanner(banner model.Banner) (int, error)
	GetUserBanner(query model.UserGet) (model.Banner, error)
	GetAllBanners(query model.AdminGet) ([]model.Banner, error) 
	UpdateBanner(banner model.PatchBanner) error
	DeleteBanner(id int) error
	FindId(id int) error
}

type Repository struct {
	Authorization
	Banner
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Banner:        NewBannerPostgres(db),
	}
}
