package routes

import (
	"HexagonalGoTemplate/src/common"
	"HexagonalGoTemplate/src/core"
	"github.com/gin-gonic/gin"
)

func CreateRoute(controller common.Controller) func(c *gin.Context) {
	return func(c *gin.Context) {
		core.GetControllerByType(controller).Execute()

		//TODO: HANDLE GENERIC ERRORS
	}
}
