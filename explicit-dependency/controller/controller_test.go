package controller

import (
	"github.com/mgreg90/dependency-injection-unit-tests/exlicit-dependency/repository"
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
	repo := repository.Repository{}
	controller := Controller{repo: repo}
	
	id := "ceovlpk3c37kch6a647g"

	if _, err := controller.GetUserByID(id); err != nil {
		t.Fatal("Unexpected error with valid xid", id)
	}
}
