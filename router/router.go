package router

import (
	"github.com/gofiber/fiber/v2"
	"urlShortener/handlers"
)

func StartRouting(app *fiber.App) {
	app.Get("/", handlers.GetRoot)                     // renders index.html with active links (if any).
	app.Post("/", handlers.PostRoot)                   // handle client url and convert to shortened url and renders result.html.
	app.Get("/s/:base62Key", handlers.GetShortenedUrl) // gets original website from shortened url and redirects to there.
	app.Get("/clear", handlers.GetClear)               // clears the database and redirects to root.
	app.Get("/d/*", handlers.DelShortenedUrl)          // deletes the shortened url from the db
}
