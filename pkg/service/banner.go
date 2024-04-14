package service

import (
	"banner_service/pkg/model"
	"banner_service/pkg/repository"
)

type BannerService struct {
	repo repository.Banner
}

func NewBannerService(repo repository.Banner) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) CreateBanner(banner model.Banner) (int, error) {
	return s.repo.CreateBanner(banner)
}

func (s *BannerService) GetUserBanner(query model.UserGet) (model.JSONB, error) {
	return s.repo.GetUserBanner(query)
}

func (s *BannerService) GetAllBanners(query model.AdminGet) ([]model.Banner, error) {
	return s.repo.GetAllBanners(query)
}

func (s *BannerService) UpdateBanner(banner model.Banner) error {
	return s.repo.UpdateBanner(banner)
}

func (s *BannerService) DeleteBanner(id int) error {
	return s.repo.DeleteBanner(id)
}

func (s *BannerService) FindId(id int) error {
	return s.repo.FindId(id)
}
