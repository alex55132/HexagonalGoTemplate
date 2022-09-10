package infrastructure

import (
	"HexagonalGoTemplate/src/domain"
	"fmt"
)

type MysqlUserRepository struct {
	domain.UserRepository
}

func (rep *MysqlUserRepository) Get(id string) {
	fmt.Println("Mec")
}
