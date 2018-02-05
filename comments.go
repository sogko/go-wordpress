package wordpress

import (
	"context"
	"fmt"
	"time"
)

// Comment represents a WordPress post comment.
type Comment struct {
	ID              int            `json:"id,omitempty"`
	AvatarURL       string         `json:"avatar_url,omitempty"`
	AvatarURLs      AvatarURLS     `json:"avatar_urls,omitempty"`
	Author          int            `json:"author,omitempty"`
	AuthorEmail     string         `json:"author_email,omitempty"`
	AuthorIP        string         `json:"author_ip,omitempty"`
	AuthorName      string         `json:"author_name,omitempty"`
	AuthorURL       string         `json:"author_url,omitempty"`
	AuthorUserAgent string         `json:"author_user_agent,omitempty"`
	Content         RenderedString `json:"content,omitempty"`
	Date            Time           `json:"date,omitempty"`
	DateGMT         Time           `json:"date_gmt,omitempty"`
	Karma           int            `json:"karma,omitempty"`
	Link            string         `json:"link,omitempty"`
	Parent          int            `json:"parent,omitempty"`
	Post            int            `json:"post,omitempty"`
	Status          string         `json:"status,omitempty"`
	Type            string         `json:"type,omitempty"`
}

// CommentsService provides access to the comment related functions in the WordPress REST API.
type CommentsService service

// CommentsListOptions are options that can be passed to List().
type CommentsListOptions struct {
	After         *time.Time `url:"after,omitempty"`
	Author        int        `url:"author,omitempty"`
	AuthorExclude []int      `url:"author_exclude,omitempty"`
	Before        *time.Time `url:"before,omitempty"`
	Exclude       []int      `url:"exclude,omitempty"`
	Include       []int      `url:"include,omitempty"`
	Parent        []int      `url:"parent,omitempty"`
	ParentExclude []int      `url:"parent_exclude,omitempty"`
	Password      string     `url:"password,omitempty"`
	Post          int        `url:"post,omitempty"`
	Search        string     `url:"search,omitempty"`
	Status        string     `url:"status,omitempty"`
	Type          string     `url:"type,omitempty"`

	ListOptions
}

// List returns a list of comments.
func (c *CommentsService) List(ctx context.Context, opts *CommentsListOptions) ([]*Comment, *Response, error) {
	u, err := addOptions("comments", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	comments := []*Comment{}
	resp, err := c.client.Do(ctx, req, &comments)
	if err != nil {
		return nil, resp, err
	}
	return comments, resp, nil
}

// Create creates a new comment.
func (c *CommentsService) Create(ctx context.Context, newComment *Comment) (*Comment, *Response, error) {
	var created Comment
	resp, err := c.client.Create(ctx, "comments", newComment, &created)
	return &created, resp, err
}

// Get returns a single comment for the given id.
func (c *CommentsService) Get(ctx context.Context, id int, params interface{}) (*Comment, *Response, error) {
	var entity Comment
	entityURL := fmt.Sprintf("comments/%v", id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}

// Update updates a single comment with the given id.
func (c *CommentsService) Update(ctx context.Context, id int, post *Comment) (*Comment, *Response, error) {
	var updated Comment
	entityURL := fmt.Sprintf("comments/%v", id)
	resp, err := c.client.Update(ctx, entityURL, post, &updated)
	return &updated, resp, err
}

// Delete removes the comment with the given id.
func (c *CommentsService) Delete(ctx context.Context, id int, params interface{}) (*Comment, *Response, error) {
	var deleted Comment
	entityURL := fmt.Sprintf("comments/%v", id)
	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)
	return &deleted, resp, err
}
