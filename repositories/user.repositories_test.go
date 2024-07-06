package repositories

import (
	"testing"

	"nugu.dev/rd-vigor/db"
)

func TestCheckEmailExists(t *testing.T) {

	store := db.NewStore()
	ur := NewUserRepository(User{}, store)

	if ur.CheckEmailExists("akdaskdjk") {
		t.Fatal("aaaa")
	}

}
