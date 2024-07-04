package model

type GormTime struct {
	CreatedAt int64 `json:"created_at" gorm:"autoUpdateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type GormIDPrimaryKey struct {
	ID uint `json:"id" gorm:"primaryKey,autoIncrement"`
}
