package cache

import (
	"banner_service/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	cli *redis.Client
}

const (
	TTLCache = 5 * time.Minute
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedis(cfg Config) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()

	return &RedisCache{cli: rdb}, err
}

func (r *RedisCache) WriteBanner(data model.Banner) error {
	var err error
	for _, tag := range data.TagIds {
		bannerKey := r.configureRedisKey(data.FeatureId, tag)
		err := r.cli.Set(context.Background(), bannerKey, data, TTLCache).Err()
		if err != nil {
			return err
		}
	}
	return err
}

func (r *RedisCache) ReadBanner(input model.UserGet) (model.Banner, error) {
	var banner model.Banner
	bannerKey := r.configureRedisKey(input.FeatureId, int64(input.TagId))
	res, err := r.cli.Get(context.Background(), bannerKey).Result()
	if err != nil {
		return banner, err
	}

	err = json.Unmarshal([]byte(res), &banner)
	if err != nil {
		return banner, err
	}
	return banner, err
}

func (r *RedisCache) configureRedisKey(featureId int, tagId int64) string {
	bannerKey := fmt.Sprintf("featureId=%d", featureId)
	bannerKey += fmt.Sprintf("tagId=%d", tagId)
	return bannerKey
}

func (r *RedisCache) IsBannerExists(key string) (bool, error) {
	notExists, err := r.cli.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return notExists == 1, nil
}
