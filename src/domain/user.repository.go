package domain

type UserRepository interface {
	Get(id string)
}
