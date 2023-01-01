package interfaces

import "github.com/mgreg90/dependency-injection-unit-tests/dependency-inversion/domain"

type Repository interface {
	GetUser(id string) (*domain.User, error)
}
