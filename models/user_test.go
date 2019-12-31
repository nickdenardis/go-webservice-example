package models

import "testing"

func TestGetUserList(t *testing.T) {
	result := GetUsers()

	if len(result) != 0 {
		t.Fatal("There should be no items in the initial user list")
	}
}

func TestAddUser(t *testing.T) {
	user := User{1, "Foo", "Bar"}

	_, err := AddUser(user)

	if err == nil {
		t.Error("User must not have an ID value")
	}

	user = User{0, "Foo", "Bar"}

	u, err := AddUser(user)

	if err != nil {
		t.Error("Could not create new user")
	}

	if u.ID == 0 {
		t.Error("User ID not assigned")
	}
}

func TestGetUserByID(t *testing.T) {
	user := User{0, "Foo", "Bar"}
	u, _ := AddUser(user)

	found, _ := GetUserByID(u.ID)

	if found.ID != u.ID {
		t.Error("User could not be found")
	}

	notFoundID := found.ID + 1
	_, err := GetUserByID(notFoundID)

	if err == nil {
		t.Error("Found invalid user")
	}
}

func TestUpdateUser(t *testing.T) {
	user := User{0, "Foo", "Bar"}
	u, _ := AddUser(user)

	u.FirstName = "Baz"

	updatedUser, _ := UpdateUser(u)

	if updatedUser.FirstName != "Baz" {
		t.Error("Value not able to be updated")
	}

	u.ID++
	_, err := UpdateUser(u)

	if err == nil {
		t.Error("Should not find invalid user")
	}
}

func TestRemoveUserById(t *testing.T) {
	user := User{0, "Foo", "Bar"}
	u, _ := AddUser(user)

	removedUser := RemoveUserById(u.ID)

	if removedUser != nil {
		t.Error("Could not remove user")
	}

	u.ID++
	err := RemoveUserById(u.ID)

	if err == nil {
		t.Error("Should not find invalid user")
	}
}
