package wordpress

import (
	"fmt"
	"net/http"
)

type PostsTerm struct {
	ID          int    `json:"id,omitempty"`
	Count       int    `json:"integer,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug,omitempty"`
	Taxonomy    string `json:"taxonomy,omitempty"`
	Parent      int    `json:"parent,omitempty"`
}

type PostsTermsCollection struct {
	client     *Client
	url        string
	parent     interface{}
	parentType string
}

func (col *PostsTermsCollection) List(taxonomy string, params interface{}) ([]PostsTerm, *http.Response, []byte, error) {
	var terms []PostsTerm
	url := fmt.Sprintf("%v/%v", col.url, taxonomy)
	resp, body, err := col.client.List(url, params, &terms)
	return terms, resp, body, err
}
func (col *PostsTermsCollection) Tag() *PostsTermsTaxonomyCollection {
	return &PostsTermsTaxonomyCollection{
		client:       col.client,
		url:          fmt.Sprintf("%v/tag", col.url),
		taxonomyBase: "tag",
	}
}
func (col *PostsTermsCollection) Category() *PostsTermsTaxonomyCollection {
	return &PostsTermsTaxonomyCollection{
		client:       col.client,
		url:          fmt.Sprintf("%v/category", col.url),
		taxonomyBase: "category",
	}
}

type PostsTermsTaxonomyCollection struct {
	client       *Client
	url          string
	taxonomyBase string
}

func (col *PostsTermsTaxonomyCollection) List(params interface{}) ([]PostsTerm, *http.Response, []byte, error) {
	var terms []PostsTerm
	resp, body, err := col.client.List(col.url, params, &terms)
	return terms, resp, body, err
}
func (col *PostsTermsTaxonomyCollection) Create(id int) (*PostsTerm, *http.Response, []byte, error) {
	var created PostsTerm
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Create(entityURL, nil, &created)
	return &created, resp, body, err
}
func (col *PostsTermsTaxonomyCollection) Get(id int, params interface{}) (*PostsTerm, *http.Response, []byte, error) {
	var entity PostsTerm
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
func (col *PostsTermsTaxonomyCollection) Delete(id int, params interface{}) (*PostsTerm, *http.Response, []byte, error) {
	var deleted PostsTerm
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, resp, body, err
}
