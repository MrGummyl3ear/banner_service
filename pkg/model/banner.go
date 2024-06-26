package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	pq "github.com/lib/pq"
)

type Banner struct {
	Id        int           `json:"-" gorm:"primaryKey"`
	TagIds    pq.Int64Array `json:"tag_ids" gorm:"Index:banner_id;type:integer[]"`
	FeatureId int           `json:"feature_id"  gorm:"Index:banner_id;"`
	Content   JSONB         `json:"content"  gorm:"type:jsonb"`
	IsActive  bool          `json:"is_active" gorm:""`
	CreatedAt time.Time     `json:"created_at" gorm:""`
	UpdatedAt time.Time     `json:"updated_at" gorm:""`
}

type PatchBanner struct {
	Id        int           `json:"-" gorm:"primaryKey"`
	TagIds    pq.Int64Array `json:"tag_ids" gorm:"uniqueIndex:banner_id;type:integer[]"`
	FeatureId int           `json:"feature_id"  gorm:"uniqueIndex:banner_id;"`
	Content   JSONB         `json:"content"  gorm:"type:jsonb"`
	IsActive  *bool         `json:"is_active" gorm:""`
	CreatedAt time.Time     `json:"created_at,omitempty" gorm:""`
	UpdatedAt time.Time     `json:"updated_at,omitempty" gorm:""`
}

// JSONB Interface for JSONB Field of yourTableName Table
type JSONB map[string]interface{}

// Value Marshal
func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan Unmarshal
func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

func (b Banner) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func (b Banner) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &b)
}
