package model

type UserGet struct {
	TagId           int  `form:"tag_id" binding:"required"`
	FeatureId       int  `form:"feature_id" binding:"required"`
	UseLastRevision bool `form:"use_last_revision"`
}

type AdminGet struct {
	TagId     int `form:"tag_id"`
	FeatureId int `form:"feature_id"`
	Limit     int `form:"limit"`
	Offset    int `form:"offset"`
}
