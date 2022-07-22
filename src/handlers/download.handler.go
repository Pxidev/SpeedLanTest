package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func DownloadHandler(c *fiber.Ctx) error {
	return c.Download("./Profile.zip")
}
