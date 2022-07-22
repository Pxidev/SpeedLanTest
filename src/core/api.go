//go:generate statik -src=../../ui/dist -include=*.* -f

package core

import (
	"SpeedLanTest/cmd/utils"
	_ "SpeedLanTest/src/core/statik"
	"SpeedLanTest/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/rakyll/statik/fs"
)

type Server struct {
	IP      string
	Port    string
	Version string
}

func (i *Server) Run() {
	statik, err := fs.New()
	if err != nil {
		utils.SendLog("error", err.Error())
	}

	engine := html.NewFileSystem(statik, ".html")

	app := fiber.New(fiber.Config{
		Prefork:               false,
		CaseSensitive:         true,
		StrictRouting:         true,
		ServerHeader:          "Pixidev v" + i.Version,
		DisableStartupMessage: true,
		AppName:               "SpeedLanTest",
		Views:                 engine,
		BodyLimit:             100 * 1024 * 1024,
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root: statik,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routes.Router(app)

	utils.SendLog("server", "Server started in port "+i.Port)

	app.Listen(":" + i.Port)

}
