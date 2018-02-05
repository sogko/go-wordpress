package wordpress

import (
	"context"
	"fmt"
)

// Term represents a WordPress page/post term.
type Term struct {
	ID          int    `json:"id,omitempty"`
	Count       int    `json:"integer,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug,omitempty"`
	Taxonomy    string `json:"taxonomy,omitempty"`
	Parent      int    `json:"parent,omitempty"`
}

// TermsService provides access to the Terms related functions in the WordPress REST API.
type TermsService service

// List returns a list of terms.
func (c *TermsService) List(ctx context.Context, taxonomy string, params interface{}) ([]*Term, *Response, error) {
	var terms []*Term
	url := fmt.Sprintf("%v/%v", "terms", taxonomy)
	resp, err := c.client.List(ctx, url, params, &terms)
	return terms, resp, err
}

// Tag returns the terms taxonomy service configured for tags.
func (c *TermsService) Tag() *TermsTaxonomyService {
	return &TermsTaxonomyService{
		client:       c.client,
		url:          fmt.Sprintf("%v/tag", "terms"),
		taxonomyBase: "tag",
	}
}

// Category returns the terms taxonomy service configured for categories.
func (c *TermsService) Category() *TermsTaxonomyService {
	return &TermsTaxonomyService{
		client:       c.client,
		url:          fmt.Sprintf("%v/category", "terms"),
		taxonomyBase: "category",
	}
}

// TermsTaxonomyService contains information about a taxonomy term.
type TermsTaxonomyService struct {
	client       *Client
	url          string
	taxonomyBase string
}

// List returns a list of terms.
func (c *TermsTaxonomyService) List(ctx context.Context, params interface{}) ([]*Term, *Response, error) {
	var terms []*Term
	resp, err := c.client.List(ctx, c.url, params, &terms)
	return terms, resp, err
}

// Create creates a new term.
func (c *TermsTaxonomyService) Create(ctx context.Context, newTerm *Term) (*Term, *Response, error) {
	var created Term
	resp, err := c.client.Create(ctx, c.url, newTerm, &created)
	return &created, resp, err
}

// Get returns a single term for the given id.
func (c *TermsTaxonomyService) Get(ctx context.Context, id int, params interface{}) (*Term, *Response, error) {
	var entity Term
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}

// Update updates a single term with the given id.
func (c *TermsTaxonomyService) Update(ctx context.Context, id int, post *Term) (*Term, *Response, error) {
	var updated Term
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Update(ctx, entityURL, post, &updated)
	return &updated, resp, err
}

// Delete removes the term with the given id.
func (c *TermsTaxonomyService) Delete(ctx context.Context, id int, params interface{}) (*Term, *Response, error) {
	var deleted Term
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)
	return &deleted, resp, err
}
