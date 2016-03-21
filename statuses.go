package wordpress

import (
	"fmt"
	"net/http"
)

type Status struct {
	Name       string `json:"name,omitempty"`
	Private    bool   `json:"private,omitempty"`
	Public     bool   `json:"public,omitempty"`
	Queryable  bool   `json:"queryable,omitempty"`
	ShowInList bool   `json:"show_in_list,omitempty"`
	Slug       string `json:"slug,omitempty"`
}

type Statuses struct {
	Publish Status `json:"publish,omitempty"`
	Future  Status `json:"future,omitempty"`
	Draft   Status `json:"draft,omitempty"`
	Pending Status `json:"pending,omitempty"`
	Private Status `json:"private,omitempty"`
}
type StatusesCollection struct {
	client *Client
	url    string
}

func (col *StatusesCollection) List(params interface{}) (*Statuses, *http.Response, []byte, error) {
	var statuses Statuses
	resp, body, err := col.client.List(col.url, params, &statuses)
	return &statuses, resp, body, err
}

func (col *StatusesCollection) Get(slug string, params interface{}) (*Status, *http.Response, []byte, error) {
	var entity Status
	entityURL := fmt.Sprintf("%v/%v", col.url, slug)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
