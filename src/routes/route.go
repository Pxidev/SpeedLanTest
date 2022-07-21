package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Router(app *fiber.App) {
	app.Get("/pxi/monitor", monitor.New())
}
