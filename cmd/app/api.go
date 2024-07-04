package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
	"sample/api/config"
	"sample/api/internal/database"
	"sample/api/internal/routes"
)

// Cmd is the command to start the API service
//
//	@title						Service API
//	@version					1.0
//	@description				This is the API documentation for the API service.
//
//	@securityDefinitions.apikey	TokenAuth
//	@in							header
//	@name						x-tg
//
//	@securityDefinitions.apiKey	JwtAuth
//	@in							header
//	@name						Authorization
var Cmd = &cobra.Command{
	Use:   "api",
	Short: "API service",
	Run: func(cmd *cobra.Command, args []string) {
		config.SetupTime()

		// Start a new fiber app
		app := fiber.New()
		// Initialize default config
		app.Use(cors.New())

		appConfig, err := config.LoadAppConfig()
		if err != nil {
			panic(err)
		}

		app.Use(func(c *fiber.Ctx) error {
			//c.Locals(config.ConfigKey, cfg)

			return c.Next()
		})

		// Connect to the Database
		database.ConnectDB(appConfig.DsnDB, appConfig.IsDebugDB, false)

		// Setup routes
		routes.SetupRouters(app)

		// Listen on PORT(env)
		if err := app.Listen(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
			panic(err)
		}
	},
}
