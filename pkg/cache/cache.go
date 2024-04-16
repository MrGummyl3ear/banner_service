package cache

import "banner_service/pkg/model"

type Cache interface {
	WriteBanner(data model.Banner) error
	ReadBanner(input model.UserGet) (model.Banner, error)
}
