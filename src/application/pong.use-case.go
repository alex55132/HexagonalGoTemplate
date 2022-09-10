package application

import (
	"HexagonalGoTemplate/src/common"
	"HexagonalGoTemplate/src/core"
	"HexagonalGoTemplate/src/domain"
	"fmt"
)

type PongUseCase struct {
	common.Executable
}

func (uc *PongUseCase) Execute() {
	fmt.Println("Execute use case pong")

	core.GetProviderByType(&domain.UserGetter{}).Execute()
}
