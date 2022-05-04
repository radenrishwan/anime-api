package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/radenrishwan/anime-api/exception"
	"log"
)

func main() {
	scrape := NewAnimeScrape()
	router := NewAnimeRouter(scrape)

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(map[string]string{
			"status": "Heyy its works",
			"author": "Raden Mohamad Rishwan",
			"github": "radenrishwan",
		})
	})

	v1 := app.Group("api/v1/anime/")

	v1.Get("/recent-anime", router.GetRecentAnime)
	v1.Get("/list-anime/:page", router.GetListAnime)
	v1.Get("/:name", router.GetAnimeInfoByName)
	v1.Get("/:name/downloads", router.GetDownloadPage)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}
