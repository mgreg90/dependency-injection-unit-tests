package controller

import (
	"github.com/mgreg90/dependency-injection-unit-tests/dependency-inversion/domain"
	"github.com/mgreg90/dependency-injection-unit-tests/dependency-inversion/interfaces"
	"github.com/rs/xid"
)

type Controller struct {
	repo interfaces.Repository
}

func (c Controller) GetUserByID(id string) (*domain.User, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}

	return c.repo.GetUser(id)
}

func validateID(id string) error {
	_, err := xid.FromString(id)

	return err
}
