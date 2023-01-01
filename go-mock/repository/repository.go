package repository

import (
	"errors"
	"github.com/mgreg90/dependency-injection-unit-tests/go-mock/domain"
)

type Repository struct{}

// We're pretending this method makes a call to the database to look up a User by ID.
// Since we're running these tests locally, it will return an error
func (r Repository) GetUser(id string) (*domain.User, error) {
	return nil, errors.New("Failed to connect to database")
}
