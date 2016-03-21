package wordpress

import (
	"fmt"
	"net/http"
)

type TypeLabels struct {
	Name            string `json:"name,omitempty"`
	SingularName    string `json:"singular_name,omitempty"`
	AddNew          string `json:"add_new,omitempty"`
	AddNewItem      string `json:"add_new_item,omitempty"`
	EditItem        string `json:"edit_item,omitempty"`
	NewItem         string `json:"new_item,omitempty"`
	ViewItem        string `json:"view_item,omitempty"`
	SearchItems     string `json:"search_items,omitempty"`
	NotFound        string `json:"not_found,omitempty"`
	NotFoundInTrash string `json:"not_found_in_trash,omitempty"`
	ParentItemColon string `json:"parent_item_colon,omitempty"`
	AllItems        string `json:"all_items,omitempty"`
	MenuName        string `json:"menu_name,omitempty"`
	NameAdminBar    string `json:"name_admin_bar,omitempty"`
}
type Type struct {
	Description  string     `json:"description,omitempty"`
	Hierarchical bool       `json:"hierarchical,omitempty"`
	Name         string     `json:"name,omitempty"`
	Slug         string     `json:"slug,omitempty"`
	Labels       TypeLabels `json:"labels,omitempty"`
}

type Types struct {
	Post       Type `json:"post,omitempty"`
	Page       Type `json:"page,omitempty"`
	Attachment Type `json:"attachment,omitempty"`
}

type TypesCollection struct {
	client *Client
	url    string
}

func (col *TypesCollection) List(params interface{}) (*Types, *http.Response, []byte, error) {
	var types Types
	resp, body, err := col.client.List(col.url, params, &types)
	return &types, resp, body, err
}

func (col *TypesCollection) Get(slug string, params interface{}) (*Type, *http.Response, []byte, error) {
	var entity Type
	entityURL := fmt.Sprintf("%v/%v", col.url, slug)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
