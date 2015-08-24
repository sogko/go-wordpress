package wordpress

import (
	"fmt"
	"log"
	"net/http"
)

type PostMeta struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}

type PostMetaCollection struct {
	client     *Client
	url        string
	parentPost *Post
}

func (col *PostMetaCollection) List(params interface{}) ([]PostMeta, *http.Response, []byte, error) {
	var meta []PostMeta
	resp, body, err := col.client.list(col.url, params, &meta)
	return meta, resp, body, err
}
func (col *PostMetaCollection) Create(new *PostMeta) (*PostMeta, *http.Response, []byte, error) {
	var created PostMeta
	resp, body, err := col.client.create(col.url, new, &created)
	return &created, resp, body, err
}
func (col *PostMetaCollection) Get(id int, params interface{}) (*PostMeta, *http.Response, []byte, error) {
	var meta PostMeta
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.get(entityURL, params, &meta)
	return &meta, resp, body, err
}
func (col *PostMetaCollection) Update(id int, meta *PostMeta) (*PostMeta, *http.Response, []byte, error) {
	var updated PostMeta
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	log.Println("URL", entityURL)
	resp, body, err := col.client.update(entityURL, meta, &updated)
	return &updated, resp, body, err
}
func (col *PostMetaCollection) Delete(id int, params interface{}) (*GeneralResponse, *http.Response, []byte, error) {
	var response GeneralResponse
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.delete(entityURL, "force=true", &response)
	return &response, resp, body, err
}
