package controller

import (
	"github.com/mgreg90/dependency-injection-unit-tests/problem/domain"
	"github.com/mgreg90/dependency-injection-unit-tests/problem/repository"
	"github.com/rs/xid"
)

func GetUserByID(id string) (*domain.User, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}

	return repository.GetUser(id)
}

func validateID(id string) error {
	_, err := xid.FromString(id)

	return err
}
