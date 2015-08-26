package wordpress

import (
	"fmt"
	"net/http"
)

type Term struct {
	ID          int    `json:"id,omitempty"`
	Count       int    `json:"integer,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug,omitempty"`
	Taxonomy    string `json:"taxonomy,omitempty"`
	Parent      int    `json:"parent,omitempty"`
}
type TermsCollection struct {
	client *Client
	url    string
}

func (col *TermsCollection) List(taxonomy string, params interface{}) ([]Term, *http.Response, []byte, error) {
	var terms []Term
	url := fmt.Sprintf("%v/%v", col.url, taxonomy)
	resp, body, err := col.client.List(url, params, &terms)
	return terms, resp, body, err
}
func (col *TermsCollection) Tag() *TermsTaxonomyCollection {
	return &TermsTaxonomyCollection{
		client:       col.client,
		url:          fmt.Sprintf("%v/tag", col.url),
		taxonomyBase: "tag",
	}
}
func (col *TermsCollection) Category() *TermsTaxonomyCollection {
	return &TermsTaxonomyCollection{
		client:       col.client,
		url:          fmt.Sprintf("%v/category", col.url),
		taxonomyBase: "category",
	}
}

type TermsTaxonomyCollection struct {
	client       *Client
	url          string
	taxonomyBase string
}

func (col *TermsTaxonomyCollection) List(params interface{}) ([]Term, *http.Response, []byte, error) {
	var terms []Term
	resp, body, err := col.client.List(col.url, params, &terms)
	return terms, resp, body, err
}
func (col *TermsTaxonomyCollection) Create(new *Term) (*Term, *http.Response, []byte, error) {
	var created Term
	resp, body, err := col.client.Create(col.url, new, &created)
	return &created, resp, body, err
}
func (col *TermsTaxonomyCollection) Get(id int, params interface{}) (*Term, *http.Response, []byte, error) {
	var entity Term
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
func (col *TermsTaxonomyCollection) Update(id int, post *Term) (*Term, *http.Response, []byte, error) {
	var updated Term
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, post, &updated)
	return &updated, resp, body, err
}
func (col *TermsTaxonomyCollection) Delete(id int, params interface{}) (*Term, *http.Response, []byte, error) {
	var deleted Term
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Delete(entityURL, params, &deleted)
	return &deleted, resp, body, err
}
