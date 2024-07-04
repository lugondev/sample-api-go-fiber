package service

import (
	"sample/api/internal/database"
	"sample/api/internal/database/model"
	"sample/api/pkg/telegram"
)

func CreateTelegramData(tgAuthData *telegram.TeleAuthData, deviceId, platform string) {
	db := database.DB
	var tg model.Telegram

	// import or update telegram data user
	db.Where("user_id = ?", tgAuthData.User.ID).First(&tg)
	if tg.ID == 0 {
		tg = model.Telegram{
			UserId:    tgAuthData.User.ID,
			FirstName: tgAuthData.User.FirstName,
			LastName:  tgAuthData.User.LastName,
			Username:  tgAuthData.User.Username,
			IsPremium: tgAuthData.User.IsPremium,
		}
		db.Create(&tg)
	} else {
		tg.FirstName = tgAuthData.User.FirstName
		tg.LastName = tgAuthData.User.LastName
		tg.Username = tgAuthData.User.Username
		tg.IsPremium = tgAuthData.User.IsPremium
		db.Save(&tg)
	}
}
