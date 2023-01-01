package controller

import (
	"testing"
)

func TestGetUserByID_InvalidXID(t *testing.T) {
	id := "something-invalid"

	if _, err := GetUserByID(id); err == nil {
		t.Fatal("Failed to return error on invalid xid", id)
	}
}

func TestGetUserByID_ValidXID(t *testing.T) {
	id := "ceovlpk3c37kch6a647g"

	if _, err := GetUserByID(id); err != nil {
		t.Fatal("Unexpected error with valid xid", id)
	}
}
