package wordpress

import (
	"context"
	"fmt"
)

// AvatarURLS returns different sizes of the users avatar.
type AvatarURLS struct {
	Size24 string `json:"24,omitempty"`
	Size48 string `json:"48,omitempty"`
	Size96 string `json:"96,omitempty"`
}

// User represents a WordPress user.
type User struct {
	ID                int                    `json:"id,omitempty"`
	AvatarURL         string                 `json:"avatar_url,omitempty"`
	AvatarURLs        AvatarURLS             `json:"avatar_urls,omitempty"`
	Capabilities      map[string]interface{} `json:"capabilities,omitempty"`
	Description       string                 `json:"description,omitempty"`
	Email             string                 `json:"email,omitempty"`
	ExtraCapabilities map[string]interface{} `json:"extra_capabilities,omitempty"`
	FirstName         string                 `json:"first_name,omitempty"`
	LastName          string                 `json:"last_name,omitempty"`
	Link              string                 `json:"link,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Nickname          string                 `json:"nickname,omitempty"`
	RegisteredDate    Time                   `json:"registered_date,omitempty"`
	Roles             []string               `json:"roles,omitempty"`
	Slug              string                 `json:"slug,omitempty"`
	URL               string                 `json:"url,omitempty"`
	Username          string                 `json:"username,omitempty"`
	Password          string                 `json:"password,omitempty"`
	Locale            string                 `json:"locale,omitempty"`
}

// UsersService provides access to the Users related functions in the WordPress REST API.
type UsersService service

// Me returns information about the currently authenticated user.
func (c *UsersService) Me(ctx context.Context, params interface{}) (*User, *Response, error) {
	url := fmt.Sprintf("%v/me", "users")
	var user User
	resp, err := c.client.Get(ctx, url, params, &user)
	return &user, resp, err
}

// List returns a list of users.
func (c *UsersService) List(ctx context.Context, opts *UserListOptions) ([]*User, *Response, error) {
	u, err := addOptions("users", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := []*User{}
	resp, err := c.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// Create creates a new user.
func (c *UsersService) Create(ctx context.Context, newUser *User) (*User, *Response, error) {
	var created User
	resp, err := c.client.Create(ctx, "users", newUser, &created)
	return &created, resp, err
}

// Get returns a single term for the given id.
func (c *UsersService) Get(ctx context.Context, id int, params interface{}) (*User, *Response, error) {
	var entity User
	entityURL := fmt.Sprintf("users/%v", id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}

// Update updates a single term with the given id.
func (c *UsersService) Update(ctx context.Context, id int, user *User) (*User, *Response, error) {
	var updated User
	entityURL := fmt.Sprintf("users/%v", id)
	resp, err := c.client.Update(ctx, entityURL, user, &updated)
	return &updated, resp, err
}

// Delete removes the term with the given id.
func (c *UsersService) Delete(ctx context.Context, id int, params interface{}) (*User, *Response, error) {
	var deleted User
	entityURL := fmt.Sprintf("users/%v", id)
	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)
	return &deleted, resp, err
}
