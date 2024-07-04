package model

type Telegram struct {
	GormTime
	GormIDPrimaryKey
	UserId    int64  `json:"user_id" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	IsPremium bool   `json:"is_premium"`
}
