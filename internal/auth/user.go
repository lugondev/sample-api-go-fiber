package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetLocalsClaims(c *fiber.Ctx) *JwtClaims {
	user := c.Locals("user")
	if user == nil {
		return nil
	}
	claims := user.(*jwt.Token).Claims.(jwt.MapClaims)
	if claims["account_id"] == "" || claims["exp"].(float64) == 0 {
		return nil
	}
	return ConvertMapClaims(claims)
}
