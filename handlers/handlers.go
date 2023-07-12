package handlers

import (
	"github.com/gofiber/fiber/v2"
	"html/template"
)

func GetRoot(c *fiber.Ctx) error {
	results, err := renderResults()
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Render("index", fiber.Map{
		"urls": results,
	})
}

func PostRoot(c *fiber.Ctx) error {
	shortUrl, err := saveURL(c.FormValue("url"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Render("result", fiber.Map{
		"url": template.URL(shortUrl),
	})
}

func GetShortenedUrl(c *fiber.Ctx) error {
	url, err := retrieveUrl(c.Params("base62Key"))
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("//" + url)
}

func GetClear(c *fiber.Ctx) error {
	if err := clearDB(); err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("/")
}

func DelShortenedUrl(c *fiber.Ctx) error {
	if err := deleteUrl(c.Params("*")); err != nil {
		return c.SendString(err.Error())
	}
	return c.Redirect("/")
}
