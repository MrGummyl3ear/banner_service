package cache

import (
	"banner_service/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
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
	bannerKey := r.configureRedisKey(data.FeatureId, data.TagIds)
	err := r.cli.Set(context.Background(), bannerKey, data, TTLCache).Err()

	return err
}

func (r *RedisCache) ReadBanner(input model.UserGet) (model.Banner, error) {
	var banner model.Banner
	bannerKey := r.configureRedisKey(input.FeatureId, pq.Int64Array{int64(input.TagId)})
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

func (r *RedisCache) configureRedisKey(featureId int, tagIds pq.Int64Array) string {
	bannerKey := fmt.Sprintf("featureId=%d", featureId)
	for i, tag := range tagIds {
		bannerKey += fmt.Sprintf("tagId%d=%d", i+1, tag)
	}
	return bannerKey
}

func (r *RedisCache) IsBannerExists(key string) (bool, error) {
	notExists, err := r.cli.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return notExists == 1, nil
}
