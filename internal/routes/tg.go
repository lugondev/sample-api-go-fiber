package routes

import (
	"github.com/gofiber/fiber/v2"
	teleHandler "sample/api/internal/routes/handler/tg"
)

func SetupTgRoutes(router fiber.Router) {
	tg := router.Group("/tg")
	tg.Post("/auth", teleHandler.RequestAuth)
}
