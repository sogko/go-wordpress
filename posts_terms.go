package wordpress

import (
	"context"
	"fmt"
)

// PostsTerm represents a WordPress post post term.
type PostsTerm struct {
	ID          int    `json:"id,omitempty"`
	Count       int    `json:"integer,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug,omitempty"`
	Taxonomy    string `json:"taxonomy,omitempty"`
	Parent      int    `json:"parent,omitempty"`
}

// PostsTermsService provides access to the post term related functions in the WordPress REST API.
type PostsTermsService struct {
	client     *Client
	url        string
	parent     interface{}
	parentType string
}

// List returns a list of post terms.
func (c *PostsTermsService) List(ctx context.Context, taxonomy string, params interface{}) ([]*PostsTerm, *Response, error) {
	var terms []*PostsTerm
	url := fmt.Sprintf("%v/%v", c.url, taxonomy)
	resp, err := c.client.List(ctx, url, params, &terms)
	return terms, resp, err
}

// Tag returns the tags of a post.
func (c *PostsTermsService) Tag() *PostsTermsTaxonomyService {
	return &PostsTermsTaxonomyService{
		client:       c.client,
		url:          fmt.Sprintf("%v/tag", c.url),
		taxonomyBase: "tag",
	}
}

// Category returns the categories of a post.
func (c *PostsTermsService) Category() *PostsTermsTaxonomyService {
	return &PostsTermsTaxonomyService{
		client:       c.client,
		url:          fmt.Sprintf("%v/category", c.url),
		taxonomyBase: "category",
	}
}

// PostsTermsTaxonomyService contains data about the post terms taxonomy service
type PostsTermsTaxonomyService struct {
	client       *Client
	url          string
	taxonomyBase string
}

// List returns a list of post terms.
func (c *PostsTermsTaxonomyService) List(ctx context.Context, params interface{}) ([]*PostsTerm, *Response, error) {
	var terms []*PostsTerm
	resp, err := c.client.List(ctx, c.url, params, &terms)
	return terms, resp, err
}

// Create creates a new post term.
func (c *PostsTermsTaxonomyService) Create(ctx context.Context, id int) (*PostsTerm, *Response, error) {
	var created PostsTerm
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Create(ctx, entityURL, nil, &created)
	return &created, resp, err
}

// Get returns a single post term for the given id.
func (c *PostsTermsTaxonomyService) Get(ctx context.Context, id int, params interface{}) (*PostsTerm, *Response, error) {
	var entity PostsTerm
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}

// Delete removes the post term with the given id.
func (c *PostsTermsTaxonomyService) Delete(ctx context.Context, id int, params interface{}) (*PostsTerm, *Response, error) {
	var deleted PostsTerm
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)
	return &deleted, resp, err
}
