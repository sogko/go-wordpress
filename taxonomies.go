package wordpress

import (
	"fmt"
	"net/http"
)

type Taxonomy struct {
	Description  string                 `json:"description,omitempty"`
	Hierarchical bool                   `json:"hierarchical,omitempty"`
	Labels       map[string]interface{} `json:"labels,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Slug         string                 `json:"slug,omitempty"`
	ShowCloud    bool                   `json:"show_cloud,omitempty"`
	Types        []string               `json:"types,omitempty"`
}
type TaxonomiesCollection struct {
	client *Client
	url    string
}

func (col *TaxonomiesCollection) List(params interface{}) (map[string]Taxonomy, *http.Response, []byte, error) {
	var taxonomies map[string]Taxonomy
	resp, body, err := col.client.List(col.url, params, &taxonomies)
	return taxonomies, resp, body, err
}

func (col *TaxonomiesCollection) Get(slug string, params interface{}) (*Taxonomy, *http.Response, []byte, error) {
	var taxonomy Taxonomy
	entityURL := fmt.Sprintf("%v/%v", col.url, slug)
	resp, body, err := col.client.Get(entityURL, params, &taxonomy)
	return &taxonomy, resp, body, err
}
