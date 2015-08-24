package wordpress

import (
	"net/http"
)

type PostRevision struct {
	ID          int    `json:"id"`
	Author      string `json:"author"` // TODO: File a WP-API bug, why am I getting string
	Date        string `json:"date"`
	DateGMT     string `json:"dateGMT"`
	GUID        string `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGMT string `json:"modifiedGMT"`
	Parent      int    `json:"parent"`
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Excerpt     string `json:"excerpt"`
}

type PostRevisionsCollection struct {
	client     *Client
	url        string
	parentPost *Post
}

func (col *PostRevisionsCollection) List(params interface{}) ([]PostRevision, *http.Response, []byte, error) {
	var revisions []PostRevision
	resp, body, err := col.client.list(col.url, params, &revisions)
	return revisions, resp, body, err
}
