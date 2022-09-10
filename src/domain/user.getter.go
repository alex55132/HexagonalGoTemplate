package domain

import "HexagonalGoTemplate/src/common"

type UserGetter struct {
	repository UserRepository
	common.Executable
}

func NewUserGetter(rep UserRepository) UserGetter {
	return UserGetter{repository: rep}
}

func (uc *UserGetter) Execute() {
	uc.repository.Get("asd")
}
