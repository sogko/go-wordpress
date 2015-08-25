package wordpress

import (
	"fmt"
	"net/http"
)

type AvatarURLS struct {
	Size24 string `json:"24"`
	Size48 string `json:"48"`
	Size96 string `json:"96"`
}
type User struct {
	ID                int                    `json:"id"`
	AvatarURL         string                 `json:"avatar_url"`
	AvatarURLs        AvatarURLS             `json:"avatar_urls"`
	Capabilities      map[string]interface{} `json:"capabilities"`
	Description       string                 `json:"description"`
	Email             string                 `json:"email"`
	ExtraCapabilities map[string]interface{} `json:"extra_capabilities"`
	FirstName         string                 `json:"first_name"`
	LastName          string                 `json:"last_name"`
	Link              string                 `json:"link"`
	Name              string                 `json:"name"`
	Nickname          string                 `json:"nickname"`
	RegisteredDate    string                 `json:"registered_date"`
	Roles             []string               `json:"roles"`
	Slug              string                 `json:"slug"`
	URL               string                 `json:"url"`
	Username          string                 `json:"username"`
	Password          string                 `json:"password"`
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