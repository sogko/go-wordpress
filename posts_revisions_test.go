package wordpress_test

import (
	"github.com/sogko/go-wordpress"
	"net/http"
	"testing"
)

func TestPostsRevisions_InvalidCall(t *testing.T) {
	// User is not allowed to call create wordpress.Post object manually to retrieve PostMetaCollection
	// A proper API call would inject the right PostMetaCollection, Client and other goodies into a post,
	// allowing user to call post.Revisions()
	invalidPost := wordpress.Post{}
	invalidRevisions := invalidPost.Revisions()
	if invalidRevisions != nil {
		t.Error("Expected revisions to be nil, %v", invalidRevisions)
	}
}

func TestPostsRevisionsList(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)

	revisions, resp, body, err := post.Revisions().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if revisions == nil {
		t.Errorf("Should not return nil revisions")
	}
}

func TestPostsRevisionsList_Lazy(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)
	postID := post.ID

	// Use Posts().Entity(postID) to retrieve revisions in one API call
	lazyRevisions, resp, body, err := wp.Posts().Entity(postID).Revisions().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if lazyRevisions == nil {
		t.Errorf("Should not return nil revisions")
	}
}
