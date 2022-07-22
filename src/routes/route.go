package routes

import (
	"SpeedLanTest/src/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Router(app *fiber.App) {
	app.Get("/pxi/monitor", monitor.New())
	app.Get("/api/download", handlers.DownloadHandler)
	app.Post("/api/upload", handlers.UploadHandler)
}
