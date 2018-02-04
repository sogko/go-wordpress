package wordpress

import (
	"context"
	"fmt"
	"time"
)

// MediaDetailsSizesItem provides details for a single media item's size.
type MediaDetailsSizesItem struct {
	File      string `json:"file,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	SourceURL string `json:"source_url,omitempty"`
}

// MediaDetailsSizes provides different sizes of the same media item.
type MediaDetailsSizes struct {
	Thumbnail MediaDetailsSizesItem `json:"thumbnail,omitempty"`
	Medium    MediaDetailsSizesItem `json:"medium,omitempty"`
	Large     MediaDetailsSizesItem `json:"large,omitempty"`
	SiteLogo  MediaDetailsSizesItem `json:"site-logo,omitempty"`
	Full      MediaDetailsSizesItem `json:"full,omitempty"`
}

// MediaDetails describes specific details about media.
type MediaDetails struct {
	Raw       string                 `json:"raw,omitempty"`
	Rendered  string                 `json:"rendered,omitempty"`
	Width     int                    `json:"width,omitempty"`
	Height    int                    `json:"height,omitempty"`
	File      string                 `json:"file,omitempty"`
	Sizes     MediaDetailsSizes      `json:"sizes,omitempty"`
	ImageMeta map[string]interface{} `json:"image_meta,omitempty"`
}

// MediaUploadOptions are options that can be passed to Create().
type MediaUploadOptions struct {
	Filename    string
	ContentType string
	Data        []byte
}

// Media represents a WordPress post media.
type Media struct {
	ID           int            `json:"id,omitempty"`
	Date         Time           `json:"date,omitempty"`
	DateGMT      Time           `json:"date_gmt,omitempty"`
	GUID         RenderedString `json:"guid,omitempty"`
	Link         string         `json:"link,omitempty"`
	Modified     Time           `json:"modified,omitempty"`
	ModifiedGMT  Time           `json:"modifiedGMT,omitempty"`
	Password     string         `json:"password,omitempty"`
	Slug         string         `json:"slug,omitempty"`
	Status       string         `json:"status,omitempty"`
	Type         string         `json:"type,omitempty"`
	Title        RenderedString `json:"title,omitempty"`
	Author       int            `json:"author,omitempty"`
	MediaStatus  string         `json:"media_status,omitempty"`
	PingStatus   string         `json:"ping_status,omitempty"`
	AltText      string         `json:"alt_text,omitempty"`
	Caption      RenderedString `json:"caption,omitempty"`
	Description  RenderedString `json:"description,omitempty"`
	MediaType    string         `json:"media_type,omitempty"`
	MediaDetails MediaDetails   `json:"media_details,omitempty"`
	Post         int            `json:"post,omitempty"`
	SourceURL    string         `json:"source_url,omitempty"`
}

// MediaService provides access to the media related functions in the WordPress REST API.
type MediaService service

// MediasListOptions are options that can be passed to List().
type MediasListOptions struct {
	After         *time.Time `url:"after,omitempty"`
	Author        []int      `url:"author,omitempty"`
	AuthorExclude []int      `url:"author_exclude,omitempty"`
	Before        *time.Time `url:"before,omitempty"`
	Exclude       []int      `url:"exclude,omitempty"`
	Include       []int      `url:"include,omitempty"`
	MediaType     string     `url:"media_type,omitempty"`
	MimeType      string     `url:"mime_type,omitempty"`
	Parent        []int      `url:"parent,omitempty"`
	ParentExclude []int      `url:"parent_exclude,omitempty"`
	Search        string     `url:"search,omitempty"`
	Slug          string     `url:"slug,omitempty"`
	Status        string     `url:"status,omitempty"`

	ListOptions
}

// List returns a list of medias.
func (c *MediaService) List(ctx context.Context, opts *MediasListOptions) ([]*Media, *Response, error) {
	u, err := addOptions("media", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	media := []*Media{}
	resp, err := c.client.Do(ctx, req, &media)
	if err != nil {
		return nil, resp, err
	}
	return media, resp, nil
}

// Create creates a new media.
func (c *MediaService) Create(ctx context.Context, options *MediaUploadOptions) (*Media, *Response, error) {
	var created Media
	resp, err := c.client.PostData(ctx, "media", options.Data, options.ContentType, options.Filename, &created)
	return &created, resp, err
}

// Get returns a single media item for the given id.
func (c *MediaService) Get(ctx context.Context, id int, params interface{}) (*Media, *Response, error) {
	var entity Media
	entityURL := fmt.Sprintf("%v/%v", "media", id)
	resp, err := c.client.Get(ctx, entityURL, params, &entity)
	return &entity, resp, err
}

// Delete removes the media item with the given id.
func (c *MediaService) Delete(ctx context.Context, id int, params interface{}) (*Media, *Response, error) {
	var deleted Media
	entityURL := fmt.Sprintf("%v/%v", "media", id)
	resp, err := c.client.Delete(ctx, entityURL, params, &deleted)
	return &deleted, resp, err
}
