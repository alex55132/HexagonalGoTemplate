package infrastructure

import (
	"HexagonalGoTemplate/src/application"
	"HexagonalGoTemplate/src/common"
)

type PingController struct {
	PingUseCase application.PingUseCase
	common.Controller
}

func (controller *PingController) Execute() {
	controller.PingUseCase.Execute()
}
