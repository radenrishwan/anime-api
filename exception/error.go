package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/radenrishwan/anime-api/model"
	"net/http"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	if e, ok := err.(NotFoundError); ok {
		return ctx.Status(e.Status).JSON(model.WebResponse{
			Code:    e.Status,
			Message: "Not Found",
			Data:    e.Error(),
		})
	}

	if e, ok := err.(InputError); ok {
		return ctx.Status(e.Status).JSON(model.WebResponse{
			Code:    e.Status,
			Message: "Bad Request",
			Data:    e.Error(),
		})
	}

	return ctx.Status(http.StatusNotFound).JSON(model.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
		Data:    "Server has been attacked by cockroach, please help me. (Server Busy)",
	})
}
