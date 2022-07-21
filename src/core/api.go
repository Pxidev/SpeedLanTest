package core

import (
	"core/cmd/utils"
	"core/src/routes"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	IP      string
	Port    string
	Version string
}

func (i *Server) Run() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(cors.AllowAll())

	utils.SendLog("server", "Server started in port "+i.Port)
	
	routes.Setup(r)

	err := r.Run(i.IP + ":" + i.Port)
	if err != nil {
		utils.SendLog("error", "El puerto "+i.Port+" no esta disponible")
		return
	}

}
