package wordpress

import (
	"fmt"
	"net/http"
)

type Labels struct {
	Raw      string `json:"raw"`
	Rendered string `json:"rendered"`
}
type Taxonomy struct {
	Description  string                 `json:"description"`
	Hierarchical bool                   `json:"hierarchical"`
	Labels       map[string]interface{} `json:"labels"`
	Name         string                 `json:"name"`
	Slug         string                 `json:"slug"`
	ShowCloud    bool                   `json:"show_cloud"`
	Types        []string               `json:"types"`
}
type TaxonomiesCollection struct {
	client *Client
	url    string
}

func (col *TaxonomiesCollection) List(params interface{}) ([]Taxonomy, *http.Response, []byte, error) {
	var taxonomies []Taxonomy
	resp, body, err := col.client.List(col.url, params, &taxonomies)
	return taxonomies, resp, body, err
}

func (col *TaxonomiesCollection) Get(slug string, params interface{}) (*Taxonomy, *http.Response, []byte, error) {
	var taxonomy Taxonomy
	entityURL := fmt.Sprintf("%v/%v", col.url, slug)
	resp, body, err := col.client.Get(entityURL, params, &taxonomy)
	return &taxonomy, resp, body, err
}
