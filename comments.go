package wordpress

import (
	"fmt"
	"net/http"
)

type Comment struct {
	ID              int        `json:"id"`
	AvatarURL       string     `json:"avatar_url"`
	AvatarURLs      AvatarURLS `json:"avatar_urls"`
	Author          int        `json:"author"`
	AuthorEmail     string     `json:"author_email"`
	AuthorIP        string     `json:"author_ip"`
	AuthorName      string     `json:"author_name"`
	AuthorURL       string     `json:"author_url"`
	AuthorUserAgent string     `json:"author_user_agent"`
	Content         Content    `json:"content"`
	Date            string     `json:"date"`
	DateGMT         string     `json:"date_gmt"`
	Karma           int        `json:"karma"`
	Link            string     `json:"link"`
	Parent          int        `json:"parent"`
	Post            int        `json:"post"`
	Status          string     `json:"status"`
	Type            string     `json:"type"`
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
