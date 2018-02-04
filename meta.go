package wordpress

import (
	"fmt"
	"log"
)

type Meta struct {
	ID    int    `json:"id,omitempty"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type MetaDeletedResponse struct {
	Message string `json:"message,omitempty"`
}

type MetaCollection struct {
	client     *Client
	url        string
	parent     interface{}
	parentType string
}

func (col *MetaCollection) List(params interface{}) ([]Meta, *Response, []byte, error) {
	var meta []Meta
	resp, body, err := col.client.List(col.url, params, &meta)
	return meta, newResponse(resp), body, err
}
func (col *MetaCollection) Create(new *Meta) (*Meta, *Response, []byte, error) {
	var created Meta
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, newResponse(resp), body, err
}
func (col *MetaCollection) Get(id int, params interface{}) (*Meta, *Response, []byte, error) {
	var meta Meta
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &meta)
	return &meta, newResponse(resp), body, err
}
func (col *MetaCollection) Update(id int, meta *Meta) (*Meta, *Response, []byte, error) {
	var updated Meta
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	log.Println("URL", entityURL)
	resp, body, err := col.client.Update(entityURL, meta, &updated)
	return &updated, newResponse(resp), body, err
}
func (col *MetaCollection) Delete(id int, params interface{}) (*MetaDeletedResponse, *Response, []byte, error) {
	var response MetaDeletedResponse
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &response)
	return &response, newResponse(resp), body, err
}
