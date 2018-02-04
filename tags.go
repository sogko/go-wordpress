package wordpress

import (
	"context"
	"fmt"
)

// Tag represents a WordPress page/post tag.
type Tag struct {
	ID          int    `json:"id,omitempty"`
	Count       int    `json:"count,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Taxonomy    string `json:"taxonomy,omitempty"`
}

// TagsService provides access to the Tag related functions in the WordPress REST API.
type TagsService service

// TagsListOptions are options that can be passed to List().
type TagsListOptions struct {
	Exclude   []int  `url:"exclude,omitempty"`
	HideEmpty bool   `url:"hide_empty,omitempty"`
	Include   []int  `url:"include,omitempty"`
	Parent    int    `url:"parent,omitempty"`
	Post      int    `url:"post,omitempty"`
	Search    string `url:"search,omitempty"`
	Slug      string `url:"slug,omitempty"`

	ListOptions
}

// List returns a list of tags.
func (c *TagsService) List(ctx context.Context, opts *TagsListOptions) ([]*Tag, *Response, error) {
	u, err := addOptions("tags", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tags := []*Tag{}
	resp, err := c.client.Do(ctx, req, &tags)
	if err != nil {
		return nil, resp, err
	}
	return tags, resp, nil
}

// Create creates a new tag.
func (c *TagsService) Create(ctx context.Context, new *Tag) (*Tag, *Response, error) {
	var created Tag
	resp, err := c.client.Create(ctx, "tags", new, &created)
	return &created, resp, err
}

// Get returns a single tag for the given id.
func (c *TagsService) Get(ctx context.Context, id int, params interface{}) (*Tag, *Response, error) {
	var entity Tag
	entityURL := fmt.Sprintf("%v/%v", "tags", id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}

// Update updates a single tag with the given id.
func (c *TagsService) Update(ctx context.Context, id int, post *Tag) (*Tag, *Response, error) {
	var updated Tag
	entityURL := fmt.Sprintf("%v/%v", "tags", id)
	resp, err := c.client.Update(ctx, entityURL, post, &updated)
	return &updated, resp, err
}

// Delete removes the tag with the given id.
func (c *TagsService) Delete(ctx context.Context, id int, params interface{}) (*Tag, *Response, error) {
	var deleted Tag
	entityURL := fmt.Sprintf("%v/%v", "tags", id)
	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)
	return &deleted, resp, err
}
