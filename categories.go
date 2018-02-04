package wordpress

import (
	"fmt"
)

type Category struct {
	ID          int    `json:"id"`
	Count       int    `json:"count"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Taxonomy    string `json:"taxonomy"`
	Parent      int    `json:"parent"`
}

type CategoriesCollection struct {
	client *Client
	url    string
}

func (col *CategoriesCollection) List(params interface{}) ([]Category, *Response, []byte, error) {
	var categories []Category
	resp, body, err := col.client.List(col.url, params, &categories)
	return categories, newResponse(resp), body, err
}
func (col *CategoriesCollection) Create(new *Category) (*Category, *Response, []byte, error) {
	var created Category
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, newResponse(resp), body, err
}
func (col *CategoriesCollection) Get(id int, params interface{}) (*Category, *Response, []byte, error) {
	var entity Category
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, newResponse(resp), body, err
}
func (col *CategoriesCollection) Update(id int, post *Category) (*Category, *Response, []byte, error) {
	var updated Category
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, post, &updated)
	return &updated, newResponse(resp), body, err
}
func (col *CategoriesCollection) Delete(id int, params interface{}) (*Category, *Response, []byte, error) {
	var deleted Category
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, newResponse(resp), body, err
}
