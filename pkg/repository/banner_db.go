package repository

import (
	"banner_service/pkg/model"

	"gorm.io/gorm"
)

type BannerPostgres struct {
	db *gorm.DB
}

func NewBannerPostgres(db *gorm.DB) *BannerPostgres {
	return &BannerPostgres{db: db}
}

func (r *BannerPostgres) CreateBanner(banner model.Banner) (int, error) {
	var err error
	tx := r.db.Begin()

	err = tx.Create(&banner).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return banner.Id, tx.Commit().Error
}

func (r *BannerPostgres) GetUserBanner(query model.UserGet) (model.Banner, error) {
	var err error
	var result model.Banner
	tx := r.db.Begin()
	err = tx.Where("? = any(tag_ids) and feature_id = ?", query.TagId, query.FeatureId).First(&result).Error
	if err != nil {
		tx.Rollback()
		return result, err
	}
	return result, tx.Commit().Error
}

func (r *BannerPostgres) GetAllBanners(query model.AdminGet) ([]model.Banner, error) {
	var result []model.Banner
	tx := r.db.Begin()
	if query.TagId != 0 {
		tx = tx.Scopes(FilterTag(query))
	}
	if query.FeatureId != 0 {
		tx = tx.Scopes(FilterFeature(query))
	}
	err := tx.Scopes(Pagination(query)).Find(&result).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return result, tx.Commit().Error
}

func (r *BannerPostgres) UpdateBanner(banner model.PatchBanner) error {
	var err error
	var input model.Banner
	tx := r.db.Begin()
	updates := make(map[string]interface{})

	if banner.TagIds != nil {
		updates["tag_ids"] = banner.TagIds
	}

	if banner.FeatureId != 0 {
		updates["feature_id"] = banner.FeatureId
	}

	if banner.Content != nil {
		updates["content"] = banner.Content
	}

	if banner.IsActive != nil {
		updates["is_active"] = banner.IsActive
	}

	updates["updated_at"] = banner.UpdatedAt
	input.Id = banner.Id

	err = tx.Model(&input).Updates(updates).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *BannerPostgres) DeleteBanner(id int) error {
	var err error
	tx := r.db.Begin()
	err = tx.Delete(&model.Banner{}, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *BannerPostgres) FindId(id int) error {
	var err error
	tx := r.db.Begin()
	err = tx.First(&model.Banner{}, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func FilterTag(params model.AdminGet) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("? = any(tag_ids)", params.TagId)
	}
}

func FilterFeature(params model.AdminGet) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("? = feature_id", params.FeatureId)
	}
}

func Pagination(params model.AdminGet) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(params.Offset).Limit(params.Limit)
	}
}
