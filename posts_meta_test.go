package wordpress_test

import (
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func cleanUpPostMeta(t *testing.T, post *wordpress.Post, metaId int) {

	// note: Need to pass in `force=true` param in order to delete post meta
	deletedMeta, resp, body, err := post.Meta().Delete(metaId, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new post: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedMeta.Message != "Deleted meta" {
		t.Errorf("Unexpected response to deleted meta: %v", deletedMeta.Message)
	}

}

func TestPostsMeta_InvalidCall(t *testing.T) {
	// User is not allowed to call create wordpress.Post object manually to retrieve PostMetaCollection
	// A proper API call would inject the right PostMetaCollection, Client and other goodies into a post,
	// allowing user to call post.Meta()
	invalidPost := wordpress.Post{}
	invalidMeta := invalidPost.Meta()
	if invalidMeta != nil {
		t.Error("Expected meta to be nil, %v", invalidMeta)
	}
}

func TestPostsMetaList_NoParams(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)

	meta, resp, body, err := post.Meta().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if meta == nil {
		t.Errorf("Should not return nil meta")
	}
}

func TestPostsMetaCreate(t *testing.T) {
	wp := initTestClient()

	// get a post
	post := getAnyOnePost(t, wp)

	// create meta for retrieved post
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, body, err := post.Meta().Create(&m)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newMeta == nil {
		t.Errorf("newMeta should not be nil")
	}
	if newMeta.Key != m.Key {
		t.Errorf("newMeta.Key should be the same, %v != %v", newMeta.Key, m.Key)
	}
	if newMeta.Value != m.Value {
		t.Errorf("newMeta.Value should be the same, %v != %v", newMeta.Value, m.Key)
	}

	// clean up
	cleanUpPostMeta(t, post, newMeta.ID)
}

func TestPostsMetaGet(t *testing.T) {
	wp := initTestClient()

	// get a post
	post := getAnyOnePost(t, wp)

	// create meta for retrieved post
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, body, err := post.Meta().Create(&m)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// get meta by id for retrieved post
	metaID := newMeta.ID
	meta, resp, body, err := post.Meta().Get(metaID, nil)
	if err != nil {
		t.Errorf("Failed to get meta: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("meta.Value should be the same, %v != %v", meta.Value, m.Value)
	}
	if meta.ID != metaID {
		t.Errorf("meta.ID should be the same, %v != %v", meta.ID, metaID)
	}
	if newMeta.Key != m.Key {
		t.Errorf("meta.Key should be the same, %v != %v", meta.Key, m.Key)
	}
	if newMeta.Value != m.Value {
		t.Errorf("meta.Value should be the same, %v != %v", meta.Value, m.Value)
	}

	// clean up
	cleanUpPostMeta(t, post, newMeta.ID)
}

func TestPostsMetaGet_Lazy(t *testing.T) {

	wp := initTestClient()

	// get a post so we can have a valid Post ID and we can create a test meta to get
	post := getAnyOnePost(t, wp)
	postID := post.ID

	// create meta for retrieved post
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, _, _ := post.Meta().Create(&m)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	metaID := newMeta.ID

	// Use Posts().Entity(postID) to retrieve meta in one API call
	lazyMeta, resp, body, err := wp.Posts().Entity(postID).Meta().Get(metaID, nil)
	if err != nil {
		t.Errorf("Failed to lazy-get meta: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("meta.Value should be the same, %v != %v", lazyMeta.Value, m.Value)
	}
	if lazyMeta.ID != metaID {
		t.Errorf("meta.ID should be the same, %v != %v", lazyMeta.ID, metaID)
	}
	if lazyMeta.Key != m.Key {
		t.Errorf("meta.Key should be the same, %v != %v", lazyMeta.Key, m.Key)
	}
	if lazyMeta.Value != m.Value {
		t.Errorf("meta.Value should be the same, %v != %v", lazyMeta.Value, m.Value)
	}
}

func TestPostsMetaUpdate(t *testing.T) {
	wp := initTestClient()

	// get a post
	post := getAnyOnePost(t, wp)

	// create meta for retrieved post
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, body, err := post.Meta().Create(&m)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// get meta by id for retrieved post
	metaID := newMeta.ID
	meta, resp, body, err := post.Meta().Get(metaID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	// update meta by id for retrieved post
	meta.Value = "newTestValue"
	updatedMeta, resp, body, err := post.Meta().Update(meta.ID, meta)
	if err != nil {
		t.Errorf("Failed to update post meta: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if updatedMeta.ID != meta.ID {
		t.Errorf("updatedMeta.ID should be the same, %v != %v", updatedMeta.ID, meta.ID)
	}
	if updatedMeta.Key != meta.Key {
		t.Errorf("updatedMeta.Key should be the same, %v != %v", updatedMeta.Key, meta.Key)
	}
	if updatedMeta.Value != meta.Value {
		t.Errorf("updatedMeta.Value should be the same, %v != %v", updatedMeta.Value, meta.Value)
	}

	// clean up
	cleanUpPostMeta(t, post, newMeta.ID)
}
