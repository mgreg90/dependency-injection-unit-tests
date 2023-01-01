package controller

import (
	"github.com/golang/mock/gomock"
	"github.com/mgreg90/dependency-injection-unit-tests/go-mock/domain"
	"github.com/mgreg90/dependency-injection-unit-tests/go-mock/mocks"
	"github.com/mgreg90/dependency-injection-unit-tests/go-mock/repository"
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

func TestGetUserByID_ValidXID(t *testing.T) {
	// GoMock setup
	ctrl := gomock.NewController(t) // this is a GoMock controller. Don't confuse with our application controller.
	defer ctrl.Finish()

	repo := mocks.NewMockRepository(ctrl)
	controller := Controller{repo: repo}

	id := "ceovlpk3c37kch6a647g"

	// Configure the return values for the call to `GetUser`
	repo.EXPECT().GetUser(id).Return(&domain.User{ID: id}, nil)

	if _, err := controller.GetUserByID(id); err != nil {
		t.Fatal("Unexpected error with valid xid", id)
	}
}
