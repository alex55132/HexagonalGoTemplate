package main

import (
	"HexagonalGoTemplate/src/core/module"
	"HexagonalGoTemplate/src/infrastructure"
	"HexagonalGoTemplate/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	module.InitAppModule()

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/ping", routes.CreateRoute(&infrastructure.PingController{}))
	r.GET("/pong", routes.CreateRoute(&infrastructure.PongController{}))

	r.Run(":8080")
}
