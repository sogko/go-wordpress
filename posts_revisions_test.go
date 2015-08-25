package wordpress_test

import (
	"fmt"
	"github.com/sogko/go-wordpress"
	"net/http"
	"testing"
	"time"
)

func getLatestRevisionForPost(t *testing.T, post *wordpress.Post) *wordpress.Revision {

	revisions, resp, _, _ := post.Revisions().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(revisions) < 1 {
		t.Fatalf("Should not return empty revisions")
	}
	// get latest revision
	revisionID := revisions[0].ID
	revision, resp, _, _ := post.Revisions().Get(revisionID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return revision
}

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

func TestPostsRevisionsGet(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)
	r := getLatestRevisionForPost(t, post)

	revisionID := r.ID

	revision, resp, body, err := post.Revisions().Get(revisionID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if revision == nil {
		t.Errorf("Should not return nil revisions")
	}
}

func TestPostsRevisionsGet_Lazy(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)
	r := getLatestRevisionForPost(t, post)

	postID := post.ID
	revisionID := r.ID

	// Use Posts().Entity(postID) to retrieve revisions in one API call
	lazyRevision, resp, body, err := wp.Posts().Entity(postID).Revisions().Get(revisionID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if lazyRevision == nil {
		t.Errorf("Should not return nil revisions")
	}
}

func TestPostsRevisionsDelete_Lazy(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)

	// Edit post to create a new revision
	// Note: wordpress would only create a new revision if there is an actual change in
	// content
	now := time.Now()
	originalTitle := post.Title.Raw
	post.Title.Raw = fmt.Sprintf("%v", now.Format("20060102150405"))
	if originalTitle == post.Title.Raw {
		t.Fatalf("Flawed test, ensure that post content is modified before an update")
	}
	updatedPost, resp, _, _ := wp.Posts().Update(post.ID, post)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	r := getLatestRevisionForPost(t, updatedPost)
	postID := updatedPost.ID
	revisionID := r.ID

	// Use Posts().Entity(postID) to delete revisions in one API call
	// Note that deleting a revision does NOT reverse the changes made in the revision
	response, resp, body, err := wp.Posts().Entity(postID).Revisions().Delete(revisionID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if response == false {
		t.Errorf("Should not return false (bool) response")
	}
}
