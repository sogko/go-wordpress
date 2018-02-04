package wordpress_test

import (
	"net/http"
	"testing"
)

func TestPostsTermsList(t *testing.T) {
	t.Skipf("Not supported anymore")
	wp, ctx := initTestClient()
	post := getAnyOnePost(t, ctx, wp)
	postID := post.ID

	terms, resp, err := wp.Posts.Entity(postID).Terms().List(ctx, "tag", nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if terms == nil {
		t.Errorf("Should not return nil terms")
	}
}
