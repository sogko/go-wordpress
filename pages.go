package wordpress

import (
	"fmt"
	"net/http"
)

type Page struct {
	collection *PagesCollection `json:"-"`

	ID            int     `json:"id,omitempty"`
	Date          string  `json:"date,omitempty"`
	DateGMT       string  `json:"date_gmt,omitempty"`
	GUID          GUID    `json:"guid,omitempty"`
	Link          string  `json:"link,omitempty"`
	Modified      string  `json:"modified,omitempty"`
	ModifiedGMT   string  `json:"modifiedGMT,omitempty"`
	Password      string  `json:"password,omitempty"`
	Slug          string  `json:"slug,omitempty"`
	Status        string  `json:"status,omitempty"`
	Type          string  `json:"type,omitempty"`
	Parent        int     `json:"parent,omitempty"`
	Title         Title   `json:"title,omitempty"`
	Content       Content `json:"content,omitempty"`
	Author        int     `json:"author,omitempty"`
	Excerpt       Excerpt `json:"excerpt,omitempty"`
	FeaturedImage int     `json:"featured_image,omitempty"`
	CommentStatus string  `json:"comment_status,omitempty"`
	PingStatus    string  `json:"ping_status,omitempty"`
	MenuOrder     int     `json:"menu_order,omitempty"`
	Template      string  `json:"template,omitempty"`
}

func (entity *Page) setCollection(col *PagesCollection) {
	entity.collection = col
}
func (entity *Page) Meta() *MetaCollection {
	if entity.collection == nil {
		// missing page.collection parent. Probably Page struct was initialized manually.
		_warning("Missing parent page collection")
		return nil
	}
	return &MetaCollection{
		client:     entity.collection.client,
		parent:     entity,
		parentType: CollectionPages,
		url:        fmt.Sprintf("%v/%v/%v", entity.collection.url, entity.ID, CollectionMeta),
	}
}
func (entity *Page) Revisions() *RevisionsCollection {
	if entity.collection == nil {
		// missing page.collection parent. Probably Page struct was initialized manually, not fetched from API
		_warning("Missing parent page collection")
		return nil
	}
	return &RevisionsCollection{
		client:     entity.collection.client,
		parent:     entity,
		parentType: CollectionPages,
		url:        fmt.Sprintf("%v/%v/%v", entity.collection.url, entity.ID, CollectionRevisions),
	}
}

func (entity *Page) Populate(params interface{}) (*Page, *http.Response, []byte, error) {
	return entity.collection.Get(entity.ID, params)
}

type PagesCollection struct {
	client    *Client
	url       string
	entityURL string
}

func (col *PagesCollection) List(params interface{}) ([]Page, *http.Response, []byte, error) {
	var pages []Page
	resp, body, err := col.client.List(col.url, params, &pages)

	// set collection object for each entity which has sub-collection
	for _, p := range pages {
		p.setCollection(col)
	}

	return pages, resp, body, err
}
func (col *PagesCollection) Create(new *Page) (*Page, *http.Response, []byte, error) {
	var created Page
	resp, body, err := col.client.Create(col.url, new, &created)

	created.setCollection(col)

	return &created, resp, body, err
}
func (col *PagesCollection) Get(id int, params interface{}) (*Page, *http.Response, []byte, error) {
	var entity Page
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)

	// set collection object for each entity which has sub-collection
	entity.setCollection(col)

	return &entity, resp, body, err
}
func (col *PagesCollection) Entity(id int) *Page {
	entity := Page{
		collection: col,
		ID:         id,
	}
	return &entity
}

func (col *PagesCollection) Update(id int, page *Page) (*Page, *http.Response, []byte, error) {
	var updated Page
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Update(entityURL, page, &updated)

	// set collection object for each entity which has sub-collection
	updated.setCollection(col)

	return &updated, resp, body, err
}
func (col *PagesCollection) Delete(id int, params interface{}) (*Page, *http.Response, []byte, error) {
	var deleted Page
	entityURL := fmt.Sprintf("%v/%v", col.url, id)

	resp, body, err := col.client.Delete(entityURL, params, &deleted)

	// set collection object for each entity which has sub-collection
	deleted.setCollection(col)

	return &deleted, resp, body, err
}
