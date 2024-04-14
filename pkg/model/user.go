package model

type User struct {
	Id       int    `json:"-" gorm:"primaryKey;unique; not null"`
	Username string `json:"username" binding:"required" gorm:"unique; not null;"`
	Password string `json:"password" binding:"required" gorm:"not null;"`
}

type Role struct {
	Username string `gorm:"uniqueIndex:not null"`
	Role     string `gorm:"uniqueIndex:not null"`
}
