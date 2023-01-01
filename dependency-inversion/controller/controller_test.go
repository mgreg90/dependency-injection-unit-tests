package controller

import (
	"github.com/mgreg90/dependency-injection-unit-tests/dependency-inversion/domain"
	"github.com/mgreg90/dependency-injection-unit-tests/dependency-inversion/repository"
	"testing"
)

func TestGetUserByID_InvalidXID(t *testing.T) {
	repo := repository.Repository{}
	controller := Controller{repo: repo}

	id := "something-invalid"

	if _, err := controller.GetUserByID(id); err == nil {
		t.Fatal("Failed to return error on invalid xid", id)
	}
}

type MockRepo struct{}

func (r MockRepo) GetUser(id string) (*domain.User, error) {
	return &domain.User{ID: id}, nil
}

func TestGetUserByID_ValidXID(t *testing.T) {
	repo := MockRepo{}
	controller := Controller{repo: repo}

	id := "ceovlpk3c37kch6a647g"

	if _, err := controller.GetUserByID(id); err != nil {
		t.Fatal("Unexpected error with valid xid", id)
	}
}
