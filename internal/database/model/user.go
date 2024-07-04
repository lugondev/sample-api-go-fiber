package model

import (
	"database/sql"
)

type User struct {
	GormTime
	GormIDPrimaryKey
	Username     string `json:"username" gorm:"unique"`
	HashPassword string `json:"hash_password"`
	PhoneNumber  sql.NullString
	Email        sql.NullString `json:"email" gorm:"unique"`
	Role         string         `json:"role"`
}
