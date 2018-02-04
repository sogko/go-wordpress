package wordpress

import (
	"context"
	"fmt"
)

// Taxonomy represents a WordPress taxonomy.
type Taxonomy struct {
	Description  string                 `json:"description,omitempty"`
	Hierarchical bool                   `json:"hierarchical,omitempty"`
	Labels       map[string]interface{} `json:"labels,omitempty"`
	Name         string                 `json:"name,omitempty"`
	ShowCloud    bool                   `json:"show_cloud,omitempty"`
	Slug         string                 `json:"slug,omitempty"`
	Types        []string               `json:"types,omitempty"`
}

// TaxonomiesService provides access to the Taxonomies related functions in the WordPress REST API.
type TaxonomiesService service

// List returns a list of taxonomies.
func (c *TaxonomiesService) List(ctx context.Context, params interface{}) (map[string]Taxonomy, *Response, error) {
	var taxonomies map[string]Taxonomy
	resp, err := c.client.List(ctx, "taxonomies", params, &taxonomies)
	return taxonomies, resp, err
}

// Get returns a single taxonomy for the given id.
func (c *TaxonomiesService) Get(ctx context.Context, slug string, params interface{}) (*Taxonomy, *Response, error) {
	var taxonomy Taxonomy
	entityURL := fmt.Sprintf("%v/%v", "taxonomies", slug)
	resp, err := c.client.Get(ctx, entityURL, params, &taxonomy)
	return &taxonomy, resp, err
}
