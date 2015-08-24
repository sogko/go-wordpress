package wordpress

import (
	"fmt"
	"net/http"
)

const (
	PostStatusDraft   = "draft"
	PostStatusPending = "pending"
	PostStatusPrivate = "private"
	PostStatusPublish = "publish"
	PostStatusTrash   = "trash"

	PostTypePost = "post"

	CommentStatusOpen   = "open"
	CommentStatusClosed = "closed"

	PingStatusOpen   = "open"
	PingStatusClosed = "closed"

	PostFormatStandard = "standard"
	PostFormatAside    = "aside"
	PostFormatGallery  = "gallery"
	PostFormatImage    = "image"
	PostFormatLink     = "link"
	PostFormatStatus   = "status"
	PostFormatQuote    = "quote"
	PostFormatVideo    = "video"
	PostFormatChat     = "chat"
)

type GUID struct {
	Raw      string `json:"raw"`
	Rendered string `json:"rendered"`
}
type Title struct {
	Raw      string `json:"raw"`
	Rendered string `json:"rendered"`
}
type Content struct {
	Raw      string `json:"raw"`
	Rendered string `json:"rendered"`
}
type Excerpt struct {
	Raw      string `json:"raw"`
	Rendered string `json:"rendered"`
}

type Post struct {
	collection *PostsCollection `json:"-"`

	ID            int     `json:"id"`
	Date          string  `json:"date"`
	DateGMT       string  `json:"date_gmt"`
	GUID          GUID    `json:"guid"`
	Link          string  `json:"link"`
	Modified      string  `json:"modified"`
	ModifiedGMT   string  `json:"modifiedGMT"`
	Password      string  `json:"password"`
	Slug          string  `json:"slug"`
	Status        string  `json:"status"`
	Type          string  `json:"type"`
	Title         Title   `json:"title"`
	Content       Content `json:"content"`
	Author        int     `json:"author"`
	Excerpt       Excerpt `json:"excerpt"`
	FeaturedImage int     `json:"featured_image"`
	CommentStatus string  `json:"comment_status"`
	PingStatus    string  `json:"ping_status"`
	Format        string  `json:"format"`
	Sticky        bool    `json:"sticky"`
}

func (entity *Post) setCollection(col *PostsCollection) {
	entity.collection = col
}
func (entity *Post) Meta() *PostMetaCollection {
	if entity.collection == nil {
		// missing post.collection parent. Probably Post struct was initialized manually.
		_warning("Missing parent post collection")
		return nil
	}
	return &PostMetaCollection{
		client:     entity.collection.client,
		parentPost: entity,
		url:        fmt.Sprintf("%v/%v/%v", entity.collection.url, entity.ID, CollectionMeta),
	}
}
func (entity *Post) Revisions() *PostRevisionsCollection {
	if entity.collection == nil {
		// missing post.collection parent. Probably Post struct was initialized manually, not fetched from API
		_warning("Missing parent post collection")
		return nil
	}
	return &PostRevisionsCollection{
		client:     entity.collection.client,
		parentPost: entity,
		url:        fmt.Sprintf("%v/%v/%v", entity.collection.url, entity.ID, CollectionRevisions),
	}
}

func (entity *Post) Populate(params interface{}) (*Post, *http.Response, []byte, error) {
	return entity.collection.Get(entity.ID, params)
}

type PostsCollection struct {
	client    *Client
	url       string
	entityURL string
}
type SamplePostParams struct {
	Filter string `url:"filter"`
}

func (col *PostsCollection) List(params interface{}) ([]Post, *http.Response, []byte, error) {
	var posts []Post
	resp, body, err := col.client.list(col.url, params, &posts)

	// set collection object for each entity which has sub-collection
	for _, p := range posts {
		p.setCollection(col)
	}

	return posts, resp, body, err
}
func (col *PostsCollection) Create(new *Post) (*Post, *http.Response, []byte, error) {
	var created Post
	resp, body, err := col.client.create(col.url, new, &created)

	created.setCollection(col)

	return &created, resp, body, err
}
func (col *PostsCollection) Get(id int, params interface{}) (*Post, *http.Response, []byte, error) {
	var entity Post
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.get(entityURL, params, &entity)

	// set collection object for each entity which has sub-collection
	entity.setCollection(col)

	return &entity, resp, body, err
}
func (col *PostsCollection) Entity(id int) *Post {
	entity := Post{
		collection: col,
		ID:         id,
	}
	return &entity
}

func (col *PostsCollection) Update(id int, post *Post) (*Post, *http.Response, []byte, error) {
	var updated Post
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.update(entityURL, post, &updated)

	// set collection object for each entity which has sub-collection
	updated.setCollection(col)

	return &updated, resp, body, err
}
func (col *PostsCollection) Delete(id int, params interface{}) (*Post, *http.Response, []byte, error) {
	var deleted Post
	entityURL := fmt.Sprintf("%v/%v", col.url, id)

	resp, body, err := col.client.delete(entityURL, params, &deleted)

	// set collection object for each entity which has sub-collection
	deleted.setCollection(col)

	return &deleted, resp, body, err
}
