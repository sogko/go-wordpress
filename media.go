package wordpress

import (
	"fmt"
	"net/http"
)

type MediaDetailsSizesItem struct {
	File      string `json:"file"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	MimeType  string `json:"mime_type"`
	SourceURL string `json:"source_url"`
}
type MediaDetailsSizes struct {
	Thumbnail MediaDetailsSizesItem `json:"thumbnail"`
	Medium    MediaDetailsSizesItem `json:"medium"`
	Large     MediaDetailsSizesItem `json:"large"`
	SiteLogo  MediaDetailsSizesItem `json:"site-logo"`
}
type MediaDetails struct {
	Raw       string                 `json:"raw"`
	Rendered  string                 `json:"rendered"`
	Width     int                    `json:"width"`
	Height    int                    `json:"height"`
	File      string                 `json:"file"`
	Sizes     MediaDetailsSizes      `json:"sizes"`
	ImageMeta map[string]interface{} `json:"image_meta"`
}
type MediaUploadOptions struct {
	Filename string
	ContentType string
	Data []byte
}
type Media struct {
	ID            int          `json:"id"`
	Date          string       `json:"date"`
	DateGMT       string       `json:"date_gmt"`
	GUID          GUID         `json:"guid"`
	Link          string       `json:"link"`
	Modified      string       `json:"modified"`
	ModifiedGMT   string       `json:"modifiedGMT"`
	Password      string       `json:"password"`
	Slug          string       `json:"slug"`
	Status        string       `json:"status"`
	Type          string       `json:"type"`
	Title         Title        `json:"title"`
	Author        int          `json:"author"`
	MediaStatus string       `json:"comment_status"`
	PingStatus    string       `json:"ping_status"`
	AltText       string       `json:"alt_text"`
	Caption       string       `json:"caption"`
	Description   string       `json:"description"`
	MediaType     string       `json:"media_type"`
	MediaDetails  MediaDetails `json:"media_details"`
	Post          int          `json:"post"`
	SourceURL     string       `json:"source_url"`
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
