package tele_handler

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"sample/api/internal/database/service"
	"sample/api/pkg/telegram"
)

// RequestAuth request authentication
//
//	@Description	Request authentication
//	@Tags			Telegram
//	@Accept			json
//	@Produce		json
//	@Param			auth	body	RequestTgAuth	true	"auth data"
//	@Success		200
//	@router			/api/tg/auth [post]
func RequestAuth(c *fiber.Ctx) error {
	tgAuth := new(RequestTgAuth)
	err := c.BodyParser(tgAuth)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	decoded, err := base64.StdEncoding.DecodeString(tgAuth.Data)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	options := telegram.ValidateOptions{ExpiresIn: telegram.AuthExpire}
	if err := telegram.Validate(string(decoded), telegram.AuthToken, options); err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	authData, _ := telegram.ParseQueryString(string(decoded))
	service.CreateTelegramData(authData, tgAuth.DeviceId, tgAuth.Platform)
	return c.JSON(fiber.Map{"status": "success", "message": "Auth success", "data": authData})
}
