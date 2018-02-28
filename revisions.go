package wordpress

import (
	"context"
	"fmt"
)

// Revision represents a WordPress page/post revision.
type Revision struct {
	ID          int            `json:"id,omitempty"`
	Author      int            `json:"author,omitempty"`
	Date        Time           `json:"date,omitempty"`
	DateGMT     Time           `json:"date_gmt,omitempty"`
	GUID        RenderedString `json:"guid,omitempty"`
	Modified    Time           `json:"modified,omitempty"`
	ModifiedGMT Time           `json:"modified_gmt,omitempty"`
	Parent      int            `json:"parent,omitempty"`
	Slug        string         `json:"slug,omitempty"`
	Title       RenderedString `json:"title,omitempty"`
	Content     RenderedString `json:"content,omitempty"`
	Excerpt     RenderedString `json:"excerpt,omitempty"`
}

// RevisionsService provides access to the revision related functions in the WordPress REST API.
type RevisionsService struct {
	service
	url        string
	parent     interface{}
	parentType string
}

// List returns a list of revisions.
func (c *RevisionsService) List(ctx context.Context, params interface{}) ([]*Revision, *Response, error) {
	var revisions []*Revision
	resp, err := c.client.List(ctx, c.url, params, &revisions)
	return revisions, resp, err
}

// Get returns a single revision for the given id.
func (c *RevisionsService) Get(ctx context.Context, id int, params interface{}) (*Revision, *Response, error) {
	var revision Revision
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Get(ctx, entityURL, params, &revision)
	return &revision, resp, err
}

// Delete removes the revision with the given id.
func (c *RevisionsService) Delete(ctx context.Context, id int, params interface{}) (*Revision, *Response, error) {
	var response Revision
	entityURL := fmt.Sprintf("%v/%v", c.url, id)
	resp, err := c.client.Delete(ctx, entityURL, "force=true", &response)
	return &response, resp, err
}
