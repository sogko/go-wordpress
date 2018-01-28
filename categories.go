package wordpress

import (
	"fmt"
	"net/http"
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

func (col *CategoriesCollection) List(params interface{}) ([]Category, *http.Response, []byte, error) {
	var categories []Category
	resp, body, err := col.client.List(col.url, params, &categories)
	return categories, resp, body, err
}
func (col *CategoriesCollection) Create(new *Category) (*Category, *http.Response, []byte, error) {
	var created Category
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, resp, body, err
}
func (col *CategoriesCollection) Get(id int, params interface{}) (*Category, *http.Response, []byte, error) {
	var entity Category
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
func (col *CategoriesCollection) Update(id int, post *Category) (*Category, *http.Response, []byte, error) {
	var updated Category
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, post, &updated)
	return &updated, resp, body, err
}
func (col *CategoriesCollection) Delete(id int, params interface{}) (*Category, *http.Response, []byte, error) {
	var deleted Category
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, resp, body, err
}
