package wordpress

import (
	"fmt"
	"net/http"
)

type Status struct {
	Name            string     `json:"name"`
	Private         bool     `json:"private"`
	Public         bool     `json:"public"`
	Queryable         bool     `json:"queryable"`
	ShowInList         bool     `json:"show_in_list"`
	Slug         string     `json:"slug"`
}

type Statuses struct {
	Publish Status `json:"publish"`
	Future Status `json:"future"`
	Draft Status `json:"draft"`
	Pending Status `json:"pending"`
	Private Status `json:"private"`
}
type StatusesCollection struct {
	client    *Client
	url       string
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
