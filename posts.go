package wordpress

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Constants for different post values.
const (
	PostStatusDraft   = "draft"
	PostStatusPending = "pending"
	PostStatusPrivate = "private"
	PostStatusPublish = "publish"
	PostStatusTrash   = "trash"

	PostTypePost = "post"
	PostTypePage = "page"

	CommentStatusOpen   = "open"
	CommentStatusClosed = "closed"

	CommentStatusApproved   = "approved"
	CommentStatusUnapproved = "unapproved"

	PingStatusOpen   = "open"
	PingStatusClosed = "closed"

	PostFormatStandard = "standard"
	PostFormatAside    = "aside"
	PostFormatGallery  = "gallery"
	PostFormatImage    = "image"
	PostFormatLink     = "link"
	PostFormatStatus   = "status"
	PostFormatQuote    = "quote"
	PostFormatVideo    = "video"
	PostFormatChat     = "chat"
)

// RenderedString contains a raw and rendered version of a string such as title, content, excerpt, etc.
type RenderedString struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}

// Post represents a WordPress post.
type Post struct {
	collection *PostsService

	Author        int            `json:"author,omitempty"`
	Categories    []int          `json:"categories,omitempty"`
	CommentStatus string         `json:"comment_status,omitempty"`
	Content       RenderedString `json:"content,omitempty"`
	Date          Time           `json:"date,omitempty"`
	DateGMT       Time           `json:"date_gmt,omitempty"`
	Excerpt       RenderedString `json:"excerpt,omitempty"`
	FeaturedMedia int            `json:"featured_media,omitempty"`
	Format        string         `json:"format,omitempty"`
	GUID          RenderedString `json:"guid,omitempty"`
	ID            int            `json:"id,omitempty"`
	Link          string         `json:"link,omitempty"`
	Modified      Time           `json:"modified,omitempty"`
	ModifiedGMT   Time           `json:"modified_gmt,omitempty"`
	Password      string         `json:"password,omitempty"`
	PingStatus    string         `json:"ping_status,omitempty"`
	Slug          string         `json:"slug,omitempty"`
	Status        string         `json:"status,omitempty"`
	Sticky        bool           `json:"sticky,omitempty"`
	Subtitle      string         `json:"wps_subtitle,omitempty"`
	Tags          []int          `json:"tags,omitempty"`
	Template      string         `json:"template,omitempty"`
	Title         RenderedString `json:"title,omitempty"`
	Type          string         `json:"type,omitempty"`
}

func (entity *Post) setService(c *PostsService) {
	entity.collection = c
}

// Revisions gets the revisions of a single post.
func (entity *Post) Revisions() *RevisionsService {
	if entity.collection == nil {
		// missing post.collection parent. Probably Post struct was initialized manually, not fetched from API
		log.Println("[go-wordpress] Missing parent post collection")
		return nil
	}
	return &RevisionsService{
		service:    service(*entity.collection),
		parent:     entity,
		parentType: "posts",
		url:        fmt.Sprintf("%v/%v/%v", "posts", entity.ID, "revisions"),
	}
}

// Terms gets the terms of a single post.
func (entity *Post) Terms() *PostsTermsService {
	if entity.collection == nil {
		// missing post.collection parent. Probably Post struct was initialized manually, not fetched from API
		log.Println("[go-wordpress] Missing parent post collection")
		return nil
	}
	return &PostsTermsService{
		client:     entity.collection.client,
		parent:     entity,
		parentType: "posts",
		url:        fmt.Sprintf("%v/%v/%v", "posts", entity.ID, "terms"),
	}
}

// Populate will fill a manually initialized post with the collection information.
func (entity *Post) Populate(ctx context.Context, params interface{}) (*Post, *Response, error) {
	return entity.collection.Get(ctx, entity.ID, params)
}

// PostsService provides access to the post related functions in the WordPress REST API.
type PostsService service

// PostsListOptions are options that can be passed to List().
type PostsListOptions struct {
	After             *time.Time `url:"after,omitempty"`                       // Limit response to posts published after a given ISO8601 compliant date.
	Author            []int      `url:"author,omitempty,brackets"`             // Limit result set to posts assigned to specific authors.
	AuthorExclude     []int      `url:"author_exclude,omitempty,brackets"`     // Ensure result set excludes posts assigned to specific authors.
	Before            *time.Time `url:"before,omitempty"`                      // Limit response to posts published before a given ISO8601 compliant date.
	Categories        []int      `url:"categories,omitempty,brackets"`         // Limit result set to posts with given category IDs.
	CategoriesExclude []int      `url:"categories_exclude,omitempty,brackets"` // Ensure result set excludes posts with given category IDs.
	Exclude           []int      `url:"exclude,omitempty,brackets"`            // Ensure result set excludes specific IDs.
	Include           []int      `url:"include,omitempty,brackets"`            // Limit result set to specific IDs.
	Search            string     `url:"search,omitempty"`                      // Limit results to those matching a string.
	Slug              string     `url:"slug,omitempty"`                        // An alphanumeric identifier for the post type.
	Status            string     `url:"status,omitempty"`                      // A named status for the object.
	Sticky            bool       `url:"sticky,omitempty"`                      // Whether or not the object should be treated as sticky.
	Tags              []int      `url:"tags,omitempty,brackets"`               // Limit result set to posts with given tag IDs.
	TagsExclude       []int      `url:"tags_exclude,omitempty,brackets"`       // Ensure result set excludes posts with given tag IDs.

	ListOptions
}

// List returns a list of posts.
func (c *PostsService) List(ctx context.Context, opts *PostsListOptions) ([]*Post, *Response, error) {
	u, err := addOptions("posts", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	posts := []*Post{}
	resp, err := c.client.Do(ctx, req, &posts)
	if err != nil {
		return nil, resp, err
	}

	// set collection object for each entity which has sub-collection
	for _, p := range posts {
		p.setService(c)
	}

	return posts, resp, nil
}

// Create creates a new post.
func (c *PostsService) Create(ctx context.Context, newPost *Post) (*Post, *Response, error) {
	var created Post
	resp, err := c.client.Create(ctx, "posts", newPost, &created)

	created.setService(c)

	return &created, resp, err
}

// Get returns a single post for the given id.
func (c *PostsService) Get(ctx context.Context, id int, params interface{}) (*Post, *Response, error) {
	var entity Post
	entityURL := fmt.Sprintf("posts/%v", id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)

	// set collection object for each entity which has sub-collection
	entity.setService(c)

	return &entity, resp, err
}

// Entity returns a basic post for the given id.
func (c *PostsService) Entity(id int) *Post {
	entity := Post{
		collection: c,
		ID:         id,
	}
	return &entity
}

// Update updates a single post with the given id.
func (c *PostsService) Update(ctx context.Context, id int, post *Post) (*Post, *Response, error) {
	var updated Post
	entityURL := fmt.Sprintf("posts/%v", id)
	resp, err := c.client.Update(ctx, entityURL, post, &updated)

	// set collection object for each entity which has sub-collection
	updated.setService(c)

	return &updated, resp, err
}

// Delete removes the post with the given id.
func (c *PostsService) Delete(ctx context.Context, id int, params interface{}) (*Post, *Response, error) {
	var deleted Post
	entityURL := fmt.Sprintf("posts/%v", id)

	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)

	// set collection object for each entity which has sub-collection
	deleted.setService(c)

	return &deleted, resp, err
}
