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
	PostTypePage = "page"

	CommentStatusOpen   = "open"
	CommentStatusClosed = "closed"

	CommentStatusApproved   = "approved"
	CommentStatusUnapproved = "unapproved"

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
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}
type Title struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}
type Content struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}
type Excerpt struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}

type Description struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}

type Caption struct {
	Raw      string `json:"raw,omitempty"`
	Rendered string `json:"rendered,omitempty"`
}

type Post struct {
	collection *PostsCollection

	Author        int     `json:"author,omitempty"`
	Categories    []int   `json:"categories,omitempty"`
	CommentStatus string  `json:"comment_status,omitempty"`
	Content       Content `json:"content,omitempty"`
	Date          Time    `json:"date,omitempty"`
	DateGMT       Time    `json:"date_gmt,omitempty"`
	Excerpt       Excerpt `json:"excerpt,omitempty"`
	FeaturedMedia int     `json:"featured_media,omitempty"`
	Format        string  `json:"format,omitempty"`
	GUID          GUID    `json:"guid,omitempty"`
	ID            int     `json:"id,omitempty"`
	Link          string  `json:"link,omitempty"`
	Modified      Time    `json:"modified,omitempty"`
	ModifiedGMT   Time    `json:"modified_gmt,omitempty"`
	ModifiedGMT   Time    `json:"modifiedGMT,omitempty"`
	Password      string  `json:"password,omitempty"`
	PingStatus    string  `json:"ping_status,omitempty"`
	Slug          string  `json:"slug,omitempty"`
	Status        string  `json:"status,omitempty"`
	Sticky        bool    `json:"sticky,omitempty"`
	Tags          []int   `json:"tags,omitempty"`
	Template      string  `json:"template,omitempty"`
	Title         Title   `json:"title,omitempty"`
	Type          string  `json:"type,omitempty"`
	WpsSubtitle   string  `json:"wps_subtitle,omitempty"`
}

func (entity *Post) setCollection(col *PostsCollection) {
	entity.collection = col
}
func (entity *Post) Meta() *MetaCollection {
	if entity.collection == nil {
		// missing post.collection parent. Probably Post struct was initialized manually.
		_warning("Missing parent post collection")
		return nil
	}
	return &MetaCollection{
		client:     entity.collection.client,
		parent:     entity,
		parentType: CollectionPosts,
		url:        fmt.Sprintf("%v/%v/%v", entity.collection.url, entity.ID, CollectionMeta),
	}
}
func (entity *Post) Revisions() *RevisionsCollection {
	if entity.collection == nil {
		// missing post.collection parent. Probably Post struct was initialized manually, not fetched from API
		_warning("Missing parent post collection")
		return nil
	}
	return &RevisionsCollection{
		client:     entity.collection.client,
		parent:     entity,
		parentType: CollectionPosts,
		url:        fmt.Sprintf("%v/%v/%v", entity.collection.url, entity.ID, CollectionRevisions),
	}
}
func (entity *Post) Terms() *PostsTermsCollection {
	if entity.collection == nil {
		// missing post.collection parent. Probably Post struct was initialized manually, not fetched from API
		_warning("Missing parent post collection")
		return nil
	}
	return &PostsTermsCollection{
		client:     entity.collection.client,
		parent:     entity,
		parentType: CollectionPosts,
		url:        fmt.Sprintf("%v/%v/%v", entity.collection.url, entity.ID, CollectionTerms),
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

func (col *PostsCollection) List(params interface{}) ([]Post, *http.Response, []byte, error) {
	var posts []Post
	resp, body, err := col.client.List(col.url, params, &posts)

	// set collection object for each entity which has sub-collection
	for _, p := range posts {
		p.setCollection(col)
	}

	return posts, resp, body, err
}
func (col *PostsCollection) Create(new *Post) (*Post, *http.Response, []byte, error) {
	var created Post
	resp, body, err := col.client.Create(col.url, new, &created)

	created.setCollection(col)

	return &created, resp, body, err
}
func (col *PostsCollection) Get(id int, params interface{}) (*Post, *http.Response, []byte, error) {
	var entity Post
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)

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
	resp, body, err := col.client.Update(entityURL, post, &updated)

	// set collection object for each entity which has sub-collection
	updated.setCollection(col)

	return &updated, resp, body, err
}
func (col *PostsCollection) Delete(id int, params interface{}) (*Post, *http.Response, []byte, error) {
	var deleted Post
	entityURL := fmt.Sprintf("%v/%v", col.url, id)

	resp, body, err := col.client.Delete(entityURL, params, &deleted)

	// set collection object for each entity which has sub-collection
	deleted.setCollection(col)

	return &deleted, resp, body, err
}
