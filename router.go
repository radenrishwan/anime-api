package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/radenrishwan/anime-api/exception"
	"github.com/radenrishwan/anime-api/model"
	"net/http"
	"strconv"
)

type animeRouter struct {
	AnimeScrape
}

func NewAnimeRouter(animeScrape AnimeScrape) AnimeRouter {
	return &animeRouter{AnimeScrape: animeScrape}
}

type AnimeRouter interface {
	GetRecentAnime(ctx *fiber.Ctx) error
	GetListAnime(ctx *fiber.Ctx) error
	GetAnimeInfoByName(ctx *fiber.Ctx) error
	GetDownloadPage(ctx *fiber.Ctx) error
	GetListGenre(ctx *fiber.Ctx) error
	GetGenre(ctx *fiber.Ctx) error
	FindAnime(ctx *fiber.Ctx) error
}

func (router *animeRouter) GetRecentAnime(ctx *fiber.Ctx) error {
	anime, err := router.AnimeScrape.GetRecentAnime()
	if err != nil {
		panic(exception.NewNotFoundError(http.StatusNotFound, err))
	}

	return ctx.JSON(model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    anime,
	})
}

func (router *animeRouter) GetListAnime(ctx *fiber.Ctx) error {
	page := ctx.Params("page", "1")
	result, err := strconv.Atoi(page)
	if err != nil {
		panic(exception.NewInputError(http.StatusBadRequest, errors.New("please input number")))
	}

	anime, err := router.AnimeScrape.GetListAnime(result)
	if err != nil {
		panic(exception.NewNotFoundError(http.StatusNotFound, err))
	}

	return ctx.JSON(model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    anime,
	})
}

func (router *animeRouter) GetAnimeInfoByName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	anime, err := router.AnimeScrape.GetAnimeInfoByName(name)
	if err != nil {
		panic(exception.NewNotFoundError(http.StatusNotFound, err))
	}

	return ctx.JSON(model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    anime,
	})
}

func (router *animeRouter) GetDownloadPage(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	anime, err := router.AnimeScrape.GetDownloadPage(name)
	if err != nil {
		panic(exception.NewNotFoundError(http.StatusNotFound, err))
	}

	return ctx.JSON(model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    anime,
	})
}

func (router *animeRouter) GetListGenre(ctx *fiber.Ctx) error {
	anime, err := router.AnimeScrape.GetListGenre()
	if err != nil {
		panic(exception.NewNotFoundError(http.StatusNotFound, err))
	}

	return ctx.JSON(model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    anime,
	})
}

func (router *animeRouter) GetGenre(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	page, err := strconv.Atoi(ctx.Params("page", "1"))
	if err != nil {
		panic(exception.NewInputError(http.StatusBadRequest, errors.New("please input number")))
	}

	anime, err := router.AnimeScrape.GetGenre(name, page)
	if err != nil {
		panic(exception.NewNotFoundError(http.StatusNotFound, err))
	}

	return ctx.JSON(model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    anime,
	})
}

func (router *animeRouter) FindAnime(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	anime, err := router.AnimeScrape.FindAnime(name)
	if err != nil {
		panic(exception.NewNotFoundError(http.StatusNotFound, err))
	}

	return ctx.JSON(model.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    anime,
	})
}
