package wordpress_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func factoryUser() *wordpress.User {
	return &wordpress.User{
		Username: "go-wordpress-test-user1",
		Name:     "go-wordpress-test-user1",
		Email:    "go-wordpress-test-user1@email.com",
		Slug:     "go-wordpress-test-user1",
		Password: "password",
	}
}

func cleanUpUser(t *testing.T, userID int) {
	wp, ctx := initTestClient()

	// Note that deleting a user requires `force=true` since `users` resource does not support trashing
	deletedUser, resp, err := wp.Users.Delete(ctx, userID, "force=true&reassign=1")
	if err != nil {
		t.Errorf("Failed to clean up new user: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedUser.ID != userID {
		t.Errorf("Deleted user ID should be the same as newly created user: %v != %v", deletedUser.ID, userID)
	}
}

func getAnyOneUser(t *testing.T, ctx context.Context, wp *wordpress.Client) *wordpress.User {

	users, resp, _ := wp.Users.List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(users) < 1 {
		t.Fatalf("Should not return empty users")
	}

	userID := users[0].ID

	user, resp, _ := wp.Users.Get(ctx, userID, "context=edit")
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	return user
}

func TestUsersList(t *testing.T) {
	client, ctx := initTestClient()

	users, resp, err := client.Users.List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if users == nil {
		t.Errorf("Should not return nil users")
	}
}

func TestUsersMe(t *testing.T) {
	wp, ctx := initTestClient()

	currentUser, resp, err := wp.Users.Me(ctx, "context=edit")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if currentUser == nil {
		t.Errorf("currentUser should not be nil")
	}
	if currentUser.Name != USER {
		t.Errorf("Logged in user has a different username: %v != %v", currentUser.Username, USER)
	}
}

func TestUsersGet(t *testing.T) {
	wp, ctx := initTestClient()

	u := getAnyOneUser(t, ctx, wp)

	user, resp, err := wp.Users.Get(ctx, u.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if user == nil {
		t.Errorf("user should not be nil")
	}
}

func TestUsersCreate(t *testing.T) {
	wp, ctx := initTestClient()

	u := &wordpress.User{
		Username: "go-wordpress-test-user1",
		Email:    "go-wordpress-test-user1@email.com",
		Name:     "go-wordpress-test-user1",
		Slug:     "go-wordpress-test-user1",
		Password: "password",
	}

	newUser, resp, err := wp.Users.Create(ctx, u)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newUser == nil {
		t.Errorf("newUser should not be nil")
	}

	// clean up
	cleanUpUser(t, newUser.ID)
}

func TestUsersDelete(t *testing.T) {
	wp, ctx := initTestClient()

	u := factoryUser()
	newUser, resp, _ := wp.Users.Create(ctx, u)
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newUser == nil {
		t.Errorf("newUser should not be nil")
	}

	// Note that deleting a user requires `force=true` since `users` resource does not support trashing
	// If not specified, a 501 NotImplemented will be returned
	deletedUser, resp, err := wp.Users.Delete(ctx, newUser.ID, "force=true&reassign=1")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedUser.ID != newUser.ID {
		t.Errorf("Deleted user ID should be the same as newly created user: %v != %v", deletedUser.ID, newUser.ID)
	}
}

func TestUsersUpdate(t *testing.T) {
	wp, ctx := initTestClient()
	u := factoryUser()

	// create user
	newUser, resp, _ := wp.Users.Create(ctx, u)
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newUser == nil {
		t.Errorf("newUser should not be nil")
	}

	// get user in `edit` context
	user, resp, _ := wp.Users.Get(ctx, newUser.ID, "context=edit")
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if user == nil {
		t.Errorf("user should not be nil")
	}

	// modify user
	newUserEmail := "testusersupdate@gmail.com"
	if user.Email == newUserEmail {
		t.Errorf("Warning: Data must be different for proper test, %v === %v", user.Email, newUserEmail)
	}
	user.Email = newUserEmail

	// update
	updatedUser, resp, err := wp.Users.Update(ctx, user.ID, user)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if updatedUser.ID != user.ID {
		t.Errorf("updatedUser ID should be the same ass specified user: %v != %v", updatedUser.ID, user.ID)
	}

	// clean up
	cleanUpUser(t, newUser.ID)
}
