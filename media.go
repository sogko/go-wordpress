package wordpress

import (
	"fmt"
	"net/http"
)

type MediaDetailsSizesItem struct {
	File      string `json:"file,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	SourceURL string `json:"source_url,omitempty"`
}
type MediaDetailsSizes struct {
	Thumbnail MediaDetailsSizesItem `json:"thumbnail,omitempty"`
	Medium    MediaDetailsSizesItem `json:"medium,omitempty"`
	Large     MediaDetailsSizesItem `json:"large,omitempty"`
	SiteLogo  MediaDetailsSizesItem `json:"site-logo,omitempty"`
}
type MediaDetails struct {
	Raw       string                 `json:"raw,omitempty"`
	Rendered  string                 `json:"rendered,omitempty"`
	Width     int                    `json:"width,omitempty"`
	Height    int                    `json:"height,omitempty"`
	File      string                 `json:"file,omitempty"`
	Sizes     MediaDetailsSizes      `json:"sizes,omitempty"`
	ImageMeta map[string]interface{} `json:"image_meta,omitempty"`
}
type MediaUploadOptions struct {
	Filename    string
	ContentType string
	Data        []byte
}
type Media struct {
	ID           int          `json:"id,omitempty"`
	Date         string       `json:"date,omitempty"`
	DateGMT      string       `json:"date_gmt,omitempty"`
	GUID         GUID         `json:"guid,omitempty"`
	Link         string       `json:"link,omitempty"`
	Modified     string       `json:"modified,omitempty"`
	ModifiedGMT  string       `json:"modifiedGMT,omitempty"`
	Password     string       `json:"password,omitempty"`
	Slug         string       `json:"slug,omitempty"`
	Status       string       `json:"status,omitempty"`
	Type         string       `json:"type,omitempty"`
	Title        Title        `json:"title,omitempty"`
	Author       int          `json:"author,omitempty"`
	MediaStatus  string       `json:"comment_status,omitempty"`
	PingStatus   string       `json:"ping_status,omitempty"`
	AltText      string       `json:"alt_text,omitempty"`
	Caption      string       `json:"caption,omitempty"`
	Description  string       `json:"description,omitempty"`
	MediaType    string       `json:"media_type,omitempty"`
	MediaDetails MediaDetails `json:"media_details,omitempty"`
	Post         int          `json:"post,omitempty"`
	SourceURL    string       `json:"source_url,omitempty"`
}
type MediaCollection struct {
	client *Client
	url    string
}

func (col *MediaCollection) List(params interface{}) ([]Media, *http.Response, []byte, error) {
	var media []Media
	resp, body, err := col.client.List(col.url, params, &media)
	return media, resp, body, err
}
func (col *MediaCollection) Create(options *MediaUploadOptions) (*Media, *http.Response, []byte, error) {
	var created Media
	resp, body, err := col.client.PostData(col.url, options.Data, options.ContentType, options.Filename, &created)
	return &created, resp, body, err
}
func (col *MediaCollection) Get(id int, params interface{}) (*Media, *http.Response, []byte, error) {
	var entity Media
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
func (col *MediaCollection) Delete(id int, params interface{}) (*Media, *http.Response, []byte, error) {
	var deleted Media
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, resp, body, err
}
