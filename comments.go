package wordpress

import (
	"fmt"
	"net/http"
)

type Comment struct {
	ID              int        `json:"id,omitempty"`
	AvatarURL       string     `json:"avatar_url,omitempty"`
	AvatarURLs      AvatarURLS `json:"avatar_urls,omitempty"`
	Author          int        `json:"author,omitempty"`
	AuthorEmail     string     `json:"author_email,omitempty"`
	AuthorIP        string     `json:"author_ip,omitempty"`
	AuthorName      string     `json:"author_name,omitempty"`
	AuthorURL       string     `json:"author_url,omitempty"`
	AuthorUserAgent string     `json:"author_user_agent,omitempty"`
	Content         Content    `json:"content,omitempty"`
	Date            string     `json:"date,omitempty"`
	DateGMT         string     `json:"date_gmt,omitempty"`
	Karma           int        `json:"karma,omitempty"`
	Link            string     `json:"link,omitempty"`
	Parent          int        `json:"parent,omitempty"`
	Post            int        `json:"post,omitempty"`
	Status          string     `json:"status,omitempty"`
	Type            string     `json:"type,omitempty"`
}

type CommentsCollection struct {
	client *Client
	url    string
}

func (col *CommentsCollection) List(params interface{}) ([]Comment, *http.Response, []byte, error) {
	var comments []Comment
	resp, body, err := col.client.List(col.url, params, &comments)
	return comments, resp, body, err
}
func (col *CommentsCollection) Create(new *Comment) (*Comment, *http.Response, []byte, error) {
	var created Comment
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, resp, body, err
}
func (col *CommentsCollection) Get(id int, params interface{}) (*Comment, *http.Response, []byte, error) {
	var entity Comment
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
func (col *CommentsCollection) Update(id int, post *Comment) (*Comment, *http.Response, []byte, error) {
	var updated Comment
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, post, &updated)
	return &updated, resp, body, err
}
func (col *CommentsCollection) Delete(id int, params interface{}) (*Comment, *http.Response, []byte, error) {
	var deleted Comment
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, resp, body, err
}
