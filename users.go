package wordpress

import (
	"fmt"
	"net/http"
)

type AvatarURLS struct {
	Size24 string `json:"24,omitempty"`
	Size48 string `json:"48,omitempty"`
	Size96 string `json:"96,omitempty"`
}
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
	RegisteredDate    string                 `json:"registered_date,omitempty"`
	Roles             []string               `json:"roles,omitempty"`
	Slug              string                 `json:"slug,omitempty"`
	URL               string                 `json:"url,omitempty"`
	Username          string                 `json:"username,omitempty"`
	Password          string                 `json:"password,omitempty"`
}

type UsersCollection struct {
	client *Client
	url    string
}

func (col *UsersCollection) Me(params interface{}) (*User, *http.Response, []byte, error) {
	url := fmt.Sprintf("%v/me", col.url)
	var user User
	resp, body, err := col.client.Get(url, params, &user)
	return &user, resp, body, err
}
func (col *UsersCollection) List(params interface{}) ([]User, *http.Response, []byte, error) {
	var users []User
	resp, body, err := col.client.List(col.url, params, &users)
	return users, resp, body, err
}
func (col *UsersCollection) Create(new *User) (*User, *http.Response, []byte, error) {
	var created User
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, resp, body, err
}
func (col *UsersCollection) Get(id int, params interface{}) (*User, *http.Response, []byte, error) {
	var entity User
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
func (col *UsersCollection) Update(id int, post *User) (*User, *http.Response, []byte, error) {
	var updated User
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, post, &updated)
	return &updated, resp, body, err
}
func (col *UsersCollection) Delete(id int, params interface{}) (*User, *http.Response, []byte, error) {
	var deleted User
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, resp, body, err
}
