package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func UploadHandler(c *fiber.Ctx) error {
	// Get first file from form field "document":

	_, err := c.FormFile("document")

	if err != nil {
		return err
	}
	// Save file to root directory:
	return c.Status(200).SendString("string()")
}
