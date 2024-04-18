package service

import (
	"banner_service/pkg/cache"
	"banner_service/pkg/model"
	"banner_service/pkg/repository"
)

type BannerService struct {
	repo  repository.Banner
	cache cache.Cache
}

func NewBannerService(repo repository.Banner, cacheInstance cache.Cache) *BannerService {
	return &BannerService{repo: repo, cache: cacheInstance}
}

func (s *BannerService) CreateBanner(banner model.Banner) (int, error) {
	return s.repo.CreateBanner(banner)
}

func (s *BannerService) GetUserBanner(query model.UserGet) (model.Banner, error) {
	if !query.UseLastRevision {
		data, err := s.cache.ReadBanner(query)
		if err == nil {
			return data, err
		}
	}

	data, err := s.repo.GetUserBanner(query)
	if err != nil {
		return data, err
	}

	go s.cache.WriteBanner(data)

	return data, err
}

func (s *BannerService) GetAllBanners(query model.AdminGet) ([]model.Banner, error) {
	return s.repo.GetAllBanners(query)
}

func (s *BannerService) UpdateBanner(banner model.PatchBanner) error {
	return s.repo.UpdateBanner(banner)
}

func (s *BannerService) DeleteBanner(id int) error {
	return s.repo.DeleteBanner(id)
}

func (s *BannerService) FindId(id int) error {
	return s.repo.FindId(id)
}
