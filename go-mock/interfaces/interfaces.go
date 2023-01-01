package interfaces

import "github.com/mgreg90/dependency-injection-unit-tests/go-mock/domain"

//go:generate mockgen -destination ../mocks/repository.go -package mocks github.com/mgreg90/dependency-injection-unit-tests/go-mock/interfaces Repository
type Repository interface {
	GetUser(id string) (*domain.User, error)
}
