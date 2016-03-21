package wordpress

import (
	"fmt"
	"net/http"
)

type Revision struct {
	ID          int    `json:"id,omitempty"`
	Author      string `json:"author,omitempty"` // TODO: File a WP-API bug, why am I getting string instead of int?
	Date        string `json:"date,omitempty"`
	DateGMT     string `json:"dateGMT,omitempty"`
	GUID        string `json:"guid,omitempty"`
	Modified    string `json:"modified,omitempty"`
	ModifiedGMT string `json:"modifiedGMT,omitempty"`
	Parent      int    `json:"parent,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Title       string `json:"title,omitempty"`
	Content     string `json:"content,omitempty"`
	Excerpt     string `json:"excerpt,omitempty"`
}

type RevisionsCollection struct {
	client     *Client
	url        string
	parent     interface{}
	parentType string
}

func (col *RevisionsCollection) List(params interface{}) ([]Revision, *http.Response, []byte, error) {
	var revisions []Revision
	resp, body, err := col.client.List(col.url, params, &revisions)
	return revisions, resp, body, err
}

func (col *RevisionsCollection) Get(id int, params interface{}) (*Revision, *http.Response, []byte, error) {
	var revision Revision
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &revision)
	return &revision, resp, body, err
}

// TODO: file an issue for inconsistent response
func (col *RevisionsCollection) Delete(id int, params interface{}) (bool, *http.Response, []byte, error) {
	var response bool
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, "force=true", &response)
	return response, resp, body, err
}
