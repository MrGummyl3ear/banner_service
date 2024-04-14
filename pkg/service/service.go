package service

import (
	"banner_service/pkg/model"
	"banner_service/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user model.User, role string) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) ([]string, error)
}

type Banner interface {
	CreateBanner(banner model.Banner) (int, error)
	GetUserBanner(query model.UserGet) (model.Banner, error)
	GetAllBanners(query model.AdminGet) ([]model.Banner,error)
	UpdateBanner(banner model.Banner) error
	DeleteBanner(id int) error
	FindId(id int) error
}

type Service struct {
	Authorization
	Banner
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Banner:        NewBannerService(repos.Banner),
	}
}
