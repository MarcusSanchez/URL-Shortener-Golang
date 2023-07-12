package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"log"
	"urlShortener/database"
	"urlShortener/router"
)

func main() {
	database.SetupMongoDB()
	defer database.ShutDownMongoDB()

	app := initializeFiber()
	router.StartRouting(app)
	startListening(app)
}

func initializeFiber() *fiber.App {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		return c.Next()
	})
	return app
}

func startListening(app *fiber.App) {
	log.Fatal(app.Listen(database.Route))
}
