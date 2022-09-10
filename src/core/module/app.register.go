package module

import (
	"HexagonalGoTemplate/src/application"
	"HexagonalGoTemplate/src/common"
	"HexagonalGoTemplate/src/core"
	"HexagonalGoTemplate/src/domain"
	"HexagonalGoTemplate/src/infrastructure"
)

func InitAppModule() {

	var controllers map[string]common.Controller = make(map[string]common.Controller)
	var providers map[string]common.Executable = make(map[string]common.Executable)
	var repositories map[string]common.Repository = make(map[string]common.Repository)

	/* Instantiate elements */

	/* Repositories */
	var userRepository infrastructure.MysqlUserRepository = infrastructure.MysqlUserRepository{}

	/* Services */
	var userGetter domain.UserGetter = domain.NewUserGetter(&userRepository)
	var pingUseCase application.PingUseCase = application.PingUseCase{}
	var pongUseCase application.PongUseCase = application.PongUseCase{}

	/* Controllers */
	var pingController infrastructure.PingController = infrastructure.PingController{PingUseCase: pingUseCase}
	var pongController infrastructure.PongController = infrastructure.PongController{PongUseCase: pongUseCase}

	/* Save elements in their arrays */

	//Repositories
	repositories[core.GetInstanceTypeName(&userRepository)] = &userRepository
	//Use cases
	providers[core.GetInstanceTypeName(&pingUseCase)] = &pingUseCase
	//Controllers
	controllers[core.GetInstanceTypeName(&pingController)] = &pingController
	controllers[core.GetInstanceTypeName(&pongController)] = &pongController
	//Services
	providers[core.GetInstanceTypeName(&userGetter)] = &userGetter

	core.NewAppModule(controllers, providers, repositories)
}
