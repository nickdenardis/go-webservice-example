package models

import "testing"

func TestGetUserList(t *testing.T) {
	// arrange

	// act
	result := GetUsers()

	// assert
	if len(result) != 0 {
		t.Fatal("There should be no items in the initial user list")
	}
}

func TestAddUser(t *testing.T) {
	// arrange
	user := User{1, "Foo", "Bar"}

	// act
	_, err := AddUser(user)

	// assert
	if err == nil {
		t.Error("User must not have an ID value")
	}
}
