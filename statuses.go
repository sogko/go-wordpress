package wordpress

import (
	"context"
	"fmt"
)

// Status represents a WordPress post status.
type Status struct {
	Name       string `json:"name,omitempty"`
	Private    bool   `json:"private,omitempty"`
	Public     bool   `json:"public,omitempty"`
	Queryable  bool   `json:"queryable,omitempty"`
	ShowInList bool   `json:"show_in_list,omitempty"`
	Slug       string `json:"slug,omitempty"`
}

// Statuses describes multiple Statuses.
type Statuses struct {
	Publish Status `json:"publish,omitempty"`
	Future  Status `json:"future,omitempty"`
	Draft   Status `json:"draft,omitempty"`
	Pending Status `json:"pending,omitempty"`
	Private Status `json:"private,omitempty"`
}

// StatusesService provides access to the Status related functions in the WordPress REST API.
type StatusesService service

// List returns a list of statuses.
func (c *StatusesService) List(ctx context.Context, params interface{}) (*Statuses, *Response, error) {
	var statuses Statuses
	resp, err := c.client.List(ctx, "statuses", params, &statuses)
	return &statuses, resp, err
}

// Get returns a single status for the given id.
func (c *StatusesService) Get(ctx context.Context, slug string, params interface{}) (*Status, *Response, error) {
	var entity Status
	entityURL := fmt.Sprintf("statuses/%v", slug)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}
