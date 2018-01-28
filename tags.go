package wordpress

import (
	"fmt"
	"net/http"
)

type Tag struct {
	ID          int    `json:"id,omitempty"`
	Count       int    `json:"count,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Taxonomy    string `json:"taxonomy,omitempty"`
}

type TagsCollection struct {
	client *Client
	url    string
}

func (col *TagsCollection) List(params interface{}) ([]Tag, *http.Response, []byte, error) {
	var tags []Tag
	resp, body, err := col.client.List(col.url, params, &tags)
	return tags, resp, body, err
}
func (col *TagsCollection) Create(new *Tag) (*Tag, *http.Response, []byte, error) {
	var created Tag
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, resp, body, err
}
func (col *TagsCollection) Get(id int, params interface{}) (*Tag, *http.Response, []byte, error) {
	var entity Tag
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
func (col *TagsCollection) Update(id int, post *Tag) (*Tag, *http.Response, []byte, error) {
	var updated Tag
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, post, &updated)
	return &updated, resp, body, err
}
func (col *TagsCollection) Delete(id int, params interface{}) (*Tag, *http.Response, []byte, error) {
	var deleted Tag
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, resp, body, err
}
