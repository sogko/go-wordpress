package wordpress_test

import (
	"github.com/sogko/go-wordpress"
	"net/http"
	"testing"
)

func factoryUser() *wordpress.User {
	return &wordpress.User{
		Username: "go-wordpress-test-user1",
		Name:     "go-wordpress-test-user1",
		Slug:     "go-wordpress-test-user1",
		Password: "password",
	}
}
func cleanUpUser(t *testing.T, userID int) {
	wp := initTestClient()

	// Note that deleting a user requires `force=true` since `users` resource does not support trashing
	deletedUser, resp, body, err := wp.Users().Delete(userID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new user: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedUser.ID != userID {
		t.Errorf("Deleted user ID should be the same as newly created user: %v != %v", deletedUser.ID, userID)
	}
}
func getAnyOneUser(t *testing.T, wp *wordpress.Client) *wordpress.User {

	users, resp, _, _ := wp.Users().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(users) < 1 {
		t.Fatalf("Should not return empty users")
	}

	userID := users[0].ID

	user, resp, _, _ := wp.Users().Get(userID, "context=edit")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	return user
}

func TestUsersList(t *testing.T) {
	client := initTestClient()

	users, resp, body, err := client.Users().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if users == nil {
		t.Errorf("Should not return nil users")
	}
}

func TestUsersMe(t *testing.T) {
	wp := initTestClient()

	currentUser, resp, body, err := wp.Users().Me("context=edit")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
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
	wp := initTestClient()

	u := getAnyOneUser(t, wp)

	user, resp, body, err := wp.Users().Get(u.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if user == nil {
		t.Errorf("user should not be nil")
	}
}

func TestUsersCreate(t *testing.T) {
	wp := initTestClient()

	u := &wordpress.User{
		Username: "go-wordpress-test-user1",
		Name:     "go-wordpress-test-user1",
		Slug:     "go-wordpress-test-user1",
		Password: "password",
	}

	newUser, resp, body, err := wp.Users().Create(u)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newUser == nil {
		t.Errorf("newUser should not be nil")
	}

	// clean up
	cleanUpUser(t, newUser.ID)
}

func TestUsersDelete(t *testing.T) {
	wp := initTestClient()

	u := factoryUser()
	newUser, resp, _, _ := wp.Users().Create(u)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newUser == nil {
		t.Errorf("newUser should not be nil")
	}

	// Note that deleting a user requires `force=true` since `users` resource does not support trashing
	// If not specified, a 501 NotImplemented will be returned
	deletedUser, resp, body, err := wp.Users().Delete(newUser.ID, "force=true")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedUser.ID != newUser.ID {
		t.Errorf("Deleted user ID should be the same as newly created user: %v != %v", deletedUser.ID, newUser.ID)
	}
}

func TestUsersUpdate(t *testing.T) {
	wp := initTestClient()
	u := factoryUser()

	// create user
	newUser, resp, _, _ := wp.Users().Create(u)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newUser == nil {
		t.Errorf("newUser should not be nil")
	}

	// get user in `edit` context
	user, resp, _, _ := wp.Users().Get(newUser.ID, "context=edit")
	if resp.StatusCode != http.StatusOK {
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
	updatedUser, resp, body, err := wp.Users().Update(user.ID, user)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if updatedUser.ID != user.ID {
		t.Errorf("updatedUser ID should be the same ass specified user: %v != %v", updatedUser.ID, user.ID)
	}

	// clean up
	cleanUpUser(t, newUser.ID)
}
