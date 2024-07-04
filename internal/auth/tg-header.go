package auth

import (
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"sample/api/pkg/telegram"
)

func HeaderTgMiddleware(c *fiber.Ctx) error {
	token := c.GetRespHeader("x-tg")
	if token == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	options := telegram.ValidateOptions{ExpiresIn: telegram.AuthExpire}
	if err := telegram.Validate(string(decoded), telegram.AuthToken, options); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return nil
}
