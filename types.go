package wordpress

import (
	"fmt"
	"net/http"
)

type TypeLabels struct {
	Name string     `json:"name"`
	SingularName string     `json:"singular_name"`
	AddNew string     `json:"add_new"`
	AddNewItem string     `json:"add_new_item"`
	EditItem string     `json:"edit_item"`
	NewItem string     `json:"new_item"`
	ViewItem string     `json:"view_item"`
	SearchItems string     `json:"search_items"`
	NotFound string     `json:"not_found"`
	NotFoundInTrash string     `json:"not_found_in_trash"`
	ParentItemColon string     `json:"parent_item_colon"`
	AllItems string     `json:"all_items"`
	MenuName string     `json:"menu_name"`
	NameAdminBar string     `json:"name_admin_bar"`
}
type Type struct {
	Description string     `json:"description"`
	Hierarchical     bool     `json:"hierarchical"`
	Name     string     `json:"name"`
	Slug     string     `json:"slug"`
	Labels     TypeLabels     `json:"labels"`
}

type Types struct {
	Post Type `json:"post"`
	Page Type `json:"page"`
	Attachment Type `json:"attachment"`
}

type TypesCollection struct {
	client    *Client
	url       string
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
