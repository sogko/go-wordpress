package wordpress

import (
	"fmt"
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

func (col *TagsCollection) List(params interface{}) ([]Tag, *Response, []byte, error) {
	var tags []Tag
	resp, body, err := col.client.List(col.url, params, &tags)
	return tags, newResponse(resp), body, err
}
func (col *TagsCollection) Create(new *Tag) (*Tag, *Response, []byte, error) {
	var created Tag
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, newResponse(resp), body, err
}
func (col *TagsCollection) Get(id int, params interface{}) (*Tag, *Response, []byte, error) {
	var entity Tag
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, newResponse(resp), body, err
}
func (col *TagsCollection) Update(id int, post *Tag) (*Tag, *Response, []byte, error) {
	var updated Tag
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, post, &updated)
	return &updated, newResponse(resp), body, err
}
func (col *TagsCollection) Delete(id int, params interface{}) (*Tag, *Response, []byte, error) {
	var deleted Tag
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, newResponse(resp), body, err
}
