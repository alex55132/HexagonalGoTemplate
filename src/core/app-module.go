package core

import (
	"HexagonalGoTemplate/src/common"
	"reflect"
)

type AppModuleStruct struct {
	controllers  map[string]common.Controller
	providers    map[string]common.Executable
	repositories map[string]common.Repository
	testmsg      string
}

var appModule *AppModuleStruct

func NewAppModule(controllers map[string]common.Controller, providers map[string]common.Executable, repositories map[string]common.Repository) *AppModuleStruct {
	appModule = &AppModuleStruct{controllers: controllers, providers: providers, repositories: repositories, testmsg: "Hola que tal"}

	return appModule
}

func GetAppModule() *AppModuleStruct {
	return appModule
}

func GetControllerByType(instance common.Controller) common.Controller {
	typeItem := reflect.TypeOf(instance).Elem().String()

	controller := appModule.GetControllers()[typeItem]

	return controller
}

func GetProviderByType(instance common.Executable) common.Executable {
	typeItem := reflect.TypeOf(instance).Elem().String()

	controller := appModule.GetProviders()[typeItem]

	return controller
}

/**
This approach is performed because of Executable interface. Since it has a method, an error is raised if tried to add the interface to the generic type

The original approach for this method was

func GetInstanceTypeName[T Controller | Repository](instance any) string {}

Since the same error was raised with the Controller interface, I've opted for a completely generic interface

*/

func GetInstanceTypeName(instance any) string {
	return reflect.TypeOf(instance).Elem().String()
}

func (am *AppModuleStruct) SetControllers(controllers map[string]common.Controller) {
	am.controllers = controllers
}

func (am *AppModuleStruct) GetControllers() map[string]common.Controller {
	return am.controllers
}

func (am *AppModuleStruct) SetProviders(providers map[string]common.Executable) {
	am.providers = providers
}

func (am *AppModuleStruct) GetProviders() map[string]common.Executable {
	return am.providers
}

func (am *AppModuleStruct) SetRepositories(repositories map[string]common.Repository) {
	am.repositories = repositories
}

func (am *AppModuleStruct) GetRepositories() map[string]common.Repository {
	return am.repositories
}

func (am *AppModuleStruct) GetTestMsg() string {
	return am.testmsg
}
