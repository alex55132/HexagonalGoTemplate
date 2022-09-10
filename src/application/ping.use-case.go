package application

import (
	"HexagonalGoTemplate/src/common"
	"HexagonalGoTemplate/src/core"
	"HexagonalGoTemplate/src/domain"
	"fmt"
)

type PingUseCase struct {
	common.Executable
}

func (uc *PingUseCase) Execute() {
	fmt.Println("Execute use case")

	core.GetProviderByType(&domain.UserGetter{}).Execute()
}
