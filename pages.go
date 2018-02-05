package wordpress

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Page represents a WordPress page.
type Page struct {
	collection *PagesService

	ID            int            `json:"id,omitempty"`
	Date          Time           `json:"date,omitempty"`
	DateGMT       Time           `json:"date_gmt,omitempty"`
	GUID          RenderedString `json:"guid,omitempty"`
	Link          string         `json:"link,omitempty"`
	Modified      Time           `json:"modified,omitempty"`
	ModifiedGMT   Time           `json:"modifiedGMT,omitempty"`
	Password      string         `json:"password,omitempty"`
	Slug          string         `json:"slug,omitempty"`
	Status        string         `json:"status,omitempty"`
	Type          string         `json:"type,omitempty"`
	Parent        int            `json:"parent,omitempty"`
	Title         RenderedString `json:"title,omitempty"`
	Content       RenderedString `json:"content,omitempty"`
	Author        int            `json:"author,omitempty"`
	Excerpt       RenderedString `json:"excerpt,omitempty"`
	FeaturedImage int            `json:"featured_image,omitempty"`
	CommentStatus string         `json:"comment_status,omitempty"`
	PingStatus    string         `json:"ping_status,omitempty"`
	MenuOrder     int            `json:"menu_order,omitempty"`
	Template      string         `json:"template,omitempty"`
}

func (entity *Page) setService(c *PagesService) {
	entity.collection = c
}

// Revisions gets the revisions of a single page.
func (entity *Page) Revisions() *RevisionsService {
	if entity.collection == nil {
		// missing page.collection parent. Probably Page struct was initialized manually, not fetched from API
		log.Println("[go-wordpress] Missing parent page collection")
		return nil
	}
	return &RevisionsService{
		service:    service(*entity.collection),
		parent:     entity,
		parentType: "pages",
		url:        fmt.Sprintf("%v/%v/%v", "pages", entity.ID, "revisions"),
	}
}

// Populate will fill a manually initialized page with the collection information.
func (entity *Page) Populate(ctx context.Context, params interface{}) (*Page, *Response, error) {
	return entity.collection.Get(ctx, entity.ID, params)
}

// PagesService provides access to the page related functions in the WordPress REST API.
type PagesService service

// PagesListOptions are options that can be passed to List().
type PagesListOptions struct {
	After             *time.Time `url:"after,omitempty"`
	Author            int        `url:"author,omitempty"`
	AuthorExclude     []int      `url:"author_exclude,omitempty"`
	Before            *time.Time `url:"before,omitempty"`
	Categories        []int      `url:"categories,omitempty"`
	CategoriesExclude []int      `url:"categories_exclude,omitempty"`
	Exclude           []int      `url:"exclude,omitempty"`
	Include           []int      `url:"include,omitempty"`
	Search            string     `url:"search,omitempty"`
	Slug              string     `url:"slug,omitempty"`
	Status            string     `url:"status,omitempty"`
	Sticky            bool       `url:"sticky,omitempty"`
	Tags              []int      `url:"tags,omitempty"`
	TagsExclude       []int      `url:"tags_exclude,omitempty"`

	ListOptions
}

// List returns a list of pages.
func (c *PagesService) List(ctx context.Context, opts *PagesListOptions) ([]*Page, *Response, error) {
	u, err := addOptions("pages", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pages := []*Page{}
	resp, err := c.client.Do(ctx, req, &pages)
	if err != nil {
		return nil, resp, err
	}

	// set collection object for each entity which has sub-collection
	for _, p := range pages {
		p.setService(c)
	}

	return pages, resp, nil
}

// Create creates a new page.
func (c *PagesService) Create(ctx context.Context, newPage *Page) (*Page, *Response, error) {
	var created Page
	resp, err := c.client.Create(ctx, "pages", newPage, &created)

	created.setService(c)

	return &created, resp, err
}

// Get returns a single page for the given id.
func (c *PagesService) Get(ctx context.Context, id int, params interface{}) (*Page, *Response, error) {
	var entity Page
	entityURL := fmt.Sprintf("pages/%v", id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)

	// set collection object for each entity which has sub-collection
	entity.setService(c)

	return &entity, resp, err
}

// Entity returns a basic page for the given id.
func (c *PagesService) Entity(id int) *Page {
	entity := Page{
		collection: c,
		ID:         id,
	}
	return &entity
}

// Update updates a single page with the given id.
func (c *PagesService) Update(ctx context.Context, id int, page *Page) (*Page, *Response, error) {
	var updated Page
	entityURL := fmt.Sprintf("pages/%v", id)
	resp, err := c.client.Update(ctx, entityURL, page, &updated)

	// set collection object for each entity which has sub-collection
	updated.setService(c)

	return &updated, resp, err
}

// Delete removes the page with the given id.
func (c *PagesService) Delete(ctx context.Context, id int, params interface{}) (*Page, *Response, error) {
	var deleted Page
	entityURL := fmt.Sprintf("pages/%v", id)

	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)

	// set collection object for each entity which has sub-collection
	deleted.setService(c)

	return &deleted, resp, err
}
