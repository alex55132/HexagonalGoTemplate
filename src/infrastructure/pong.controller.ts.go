package infrastructure

import (
	"HexagonalGoTemplate/src/application"
	"HexagonalGoTemplate/src/common"
)

type PongController struct {
	PongUseCase application.PongUseCase
	common.Controller
}

func (controller *PongController) Execute() {
	controller.PongUseCase.Execute()
}
