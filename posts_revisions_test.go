package wordpress_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/robbiet480/go-wordpress"
)

func getLatestRevisionForPost(t *testing.T, ctx context.Context, post *wordpress.Post) *wordpress.Revision {

	revisions, resp, _ := post.Revisions().List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(revisions) < 1 {
		t.Fatalf("Should not return empty revisions")
	}
	// get latest revision
	revisionID := revisions[0].ID
	revision, resp, _ := post.Revisions().Get(ctx, revisionID, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return revision
}

func TestPostsRevisions_InvalidCall(t *testing.T) {
	// User is not allowed to call create wordpress.Post object manually to retrieve PostMetaService
	// A proper API call would inject the right PostMetaService, Client and other goodies into a post,
	// allowing user to call post.Revisions()
	invalidPost := wordpress.Post{}
	invalidRevisions := invalidPost.Revisions()
	if invalidRevisions != nil {
		t.Errorf("Expected revisions to be nil, %v", invalidRevisions)
	}
}

func TestPostsRevisionsList(t *testing.T) {
	wp, ctx := initTestClient()

	post := getAnyOnePost(t, ctx, wp)

	revisions, resp, err := post.Revisions().List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if revisions == nil {
		t.Errorf("Should not return nil revisions")
	}
}

func TestPostsRevisionsList_Lazy(t *testing.T) {
	wp, ctx := initTestClient()

	post := getAnyOnePost(t, ctx, wp)
	postID := post.ID

	// Use Posts.Entity(postID) to retrieve revisions in one API call
	lazyRevisions, resp, err := wp.Posts.Entity(postID).Revisions().List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if lazyRevisions == nil {
		t.Errorf("Should not return nil revisions")
	}
}

func TestPostsRevisionsGet(t *testing.T) {
	wp, ctx := initTestClient()

	post := getAnyOnePost(t, ctx, wp)
	r := getLatestRevisionForPost(t, ctx, post)

	revisionID := r.ID

	revision, resp, err := post.Revisions().Get(ctx, revisionID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if revision == nil {
		t.Errorf("Should not return nil revisions")
	}
}

func TestPostsRevisionsGet_Lazy(t *testing.T) {
	wp, ctx := initTestClient()

	post := getAnyOnePost(t, ctx, wp)
	r := getLatestRevisionForPost(t, ctx, post)

	postID := post.ID
	revisionID := r.ID

	// Use Posts.Entity(postID) to retrieve revisions in one API call
	lazyRevision, resp, err := wp.Posts.Entity(postID).Revisions().Get(ctx, revisionID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if lazyRevision == nil {
		t.Errorf("Should not return nil revisions")
	}
}

func TestPostsRevisionsDelete_Lazy(t *testing.T) {
	wp, ctx := initTestClient()

	post := getAnyOnePost(t, ctx, wp)

	// Edit post to create a new revision
	// Note: wordpress would only create a new revision if there is an actual change in
	// content
	now := time.Now()
	originalTitle := post.Title.Raw
	post.Title.Raw = fmt.Sprintf("%v", now.Format("20060102150405"))
	if originalTitle == post.Title.Raw {
		t.Fatalf("Flawed test, ensure that post content is modified before an update")
	}
	updatedPost, resp, _ := wp.Posts.Update(ctx, post.ID, post)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	r := getLatestRevisionForPost(t, ctx, updatedPost)
	postID := updatedPost.ID
	revisionID := r.ID

	// Use Posts.Entity(postID) to delete revisions in one API call
	// Note that deleting a revision does NOT reverse the changes made in the revision
	response, resp, err := wp.Posts.Entity(postID).Revisions().Delete(ctx, revisionID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if response == nil {
		t.Errorf("Should not return nil response")
	}
}
