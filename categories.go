package wordpress

import (
	"context"
	"fmt"
)

// Category represents a WordPress post/page category.
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

// CategoriesService provides access to the category related functions in the WordPress REST API.
type CategoriesService service

// CategoriesListOptions are options that can be passed to List().
type CategoriesListOptions struct {
	Exclude []int  `url:"exclude,omitempty"`
	Include []int  `url:"include,omitempty"`
	Parent  int    `url:"parent,omitempty"`
	Post    int    `url:"post,omitempty"`
	Search  string `url:"search,omitempty"`
	Slug    string `url:"slug,omitempty"`

	ListOptions
}

// List returns a list of categories.
func (c *CategoriesService) List(ctx context.Context, opts *CategoriesListOptions) ([]*Category, *Response, error) {
	u, err := addOptions("categories", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	categories := []*Category{}
	resp, err := c.client.Do(ctx, req, &categories)
	if err != nil {
		return nil, resp, err
	}
	return categories, resp, nil
}

// Create creates a new category.
func (c *CategoriesService) Create(ctx context.Context, newCategory *Category) (*Category, *Response, error) {
	var created Category
	resp, err := c.client.Create(ctx, "categories", newCategory, &created)
	return &created, resp, err
}

// Get returns a single category for the given id.
func (c *CategoriesService) Get(ctx context.Context, id int, params interface{}) (*Category, *Response, error) {
	var entity Category
	entityURL := fmt.Sprintf("%v/%v", "categories", id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}

// Update updates a single category with the given id.
func (c *CategoriesService) Update(ctx context.Context, id int, post *Category) (*Category, *Response, error) {
	var updated Category
	entityURL := fmt.Sprintf("%v/%v", "categories", id)
	resp, err := c.client.Update(ctx, entityURL, post, &updated)
	return &updated, resp, err
}

// Delete removes the category with the given id.
func (c *CategoriesService) Delete(ctx context.Context, id int, params interface{}) (*Category, *Response, error) {
	var deleted Category
	entityURL := fmt.Sprintf("%v/%v", "categories", id)
	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)
	return &deleted, resp, err
}
