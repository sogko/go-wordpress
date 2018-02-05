package wordpress_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func factoryPost() wordpress.Post {
	return wordpress.Post{
		Title: wordpress.RenderedString{
			Raw: "TestPostsCreate",
		},
		Content: wordpress.RenderedString{
			Raw: "<h1>HEADER</h1><p>Paragraph</p>",
		},
		Excerpt: wordpress.RenderedString{
			Raw: "<h1>HEADER</h1><p>Paragraph</p>",
		},
		Format: wordpress.PostFormatImage,
		Type:   wordpress.PostTypePost,
		Status: wordpress.PostStatusDraft,
		Slug:   "test-posts-create",
		Author: 1,
	}
}

func cleanUpPost(t *testing.T, postID int) {

	wp, ctx := initTestClient()
	deletedPost, resp, err := wp.Posts.Delete(ctx, postID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new post: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedPost.ID != postID {
		t.Errorf("Deleted post ID should be the same as newly created post: %v != %v", deletedPost.ID, postID)
	}
}

func getAnyOnePost(t *testing.T, ctx context.Context, wp *wordpress.Client) *wordpress.Post {

	posts, resp, err := wp.Posts.List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(posts) < 1 {
		log.Print(err)
		log.Print(resp)
		t.Fatalf("Should not return empty posts")
	}

	postID := posts[0].ID

	post, resp, _ := wp.Posts.Get(ctx, postID, "context=edit")
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	return post
}

func TestPostsList_NoParams(t *testing.T) {
	wp, ctx := initTestClient()

	posts, resp, err := wp.Posts.List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if posts == nil {
		t.Errorf("Should not return nil posts")
	}
	if len(posts) == 0 {
		t.Errorf("Should not return empty posts")
	}
}
func TestPostsList_WithParamsString(t *testing.T) {
	wp, ctx := initTestClient()

	// assumes that API user authenticated with `edit_posts`
	posts, resp, err := wp.Posts.List(ctx, &wordpress.PostListOptions{Status: []string{"draft"}})
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if len(posts) != 0 {
		t.Errorf("Should return zero draft posts, returned %v", len(posts))
	}
	posts, resp, err = wp.Posts.List(ctx, &wordpress.PostListOptions{Status: []string{"publish"}})
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if len(posts) == 0 {
		t.Errorf("Should return at least one published posts")
	}
}

func TestPostsGet_PostExists(t *testing.T) {
	wp, ctx := initTestClient()

	post := getAnyOnePost(t, ctx, wp)
	postID := post.ID

	post, resp, err := wp.Posts.Get(ctx, postID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if post.ID != postID {
		t.Errorf("Returned post should have the same ID as specified in Get(), %v != %v", post.ID, postID)
	}
}
func TestPostsGet_PostDoesNotExists(t *testing.T) {
	wp, ctx := initTestClient()

	postID := -1

	_, resp, err := wp.Posts.Get(ctx, postID, nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp != nil && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 400 NotFound, got %v", resp.Status)
	}
}
func TestPostsGet_Lazy(t *testing.T) {
	wp, ctx := initTestClient()

	post := getAnyOnePost(t, ctx, wp)
	postID := post.ID

	//The proper way to get lazy-fetch posts. Posts.Entity() won't make any HTTP request
	lazyPost := wp.Posts.Entity(postID)
	if lazyPost == nil {
		t.Errorf("lazyPost should not be nil")
	}
	if lazyPost.ID != postID {
		t.Errorf("lazyPost should have specified ID, %v != %v", lazyPost.ID, postID)
	}
	if lazyPost.GUID.Rendered != "" {
		t.Errorf("lazyPost should not have populated GUID field, %v", lazyPost.GUID.Rendered)
	}

	// populate Post Entity
	post, resp, err := lazyPost.Populate(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if post.ID != postID {
		t.Errorf("Returned post should have the same ID as specified in Get(), %v != %v", post.ID, postID)
	}
	if post.GUID.Rendered == "" {
		t.Errorf("post should have populated GUID field, %v", lazyPost.GUID.Rendered)
	}
}

func TestPostsCreate(t *testing.T) {
	wp, ctx := initTestClient()

	p := factoryPost()
	newPost, resp, err := wp.Posts.Create(ctx, &p)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newPost == nil {
		t.Errorf("newPost should not be nil")
	}
	if newPost.ID < 1 {
		t.Errorf("newPost.ID should not be invalid")
	}
	if newPost.Title.Raw != p.Title.Raw {
		t.Errorf("newPost.Title should be the same, %v != %v", newPost.Title.Raw, p.Title.Raw)
	}
	if newPost.Format != p.Format {
		t.Errorf("newPost.Format should be the same, %v != %v", newPost.Format, p.Format)
	}
	if newPost.Status != p.Status {
		t.Errorf("newPost.Status should be the same, %v != %v", newPost.Status, p.Status)
	}
	if newPost.Slug != p.Slug {
		t.Errorf("newPost.Slug should be the same, %v != %v", newPost.Slug, p.Slug)
	}

	// clean up
	cleanUpPost(t, newPost.ID)
}

func TestPostsUpdate(t *testing.T) {
	wp, ctx := initTestClient()

	// create a new post first
	p := factoryPost()
	newPost, resp, _ := wp.Posts.Create(ctx, &p)
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %v", resp.Status)
	}

	// get the post in `edit` context
	post, resp, _ := wp.Posts.Get(ctx, newPost.ID, "context=edit")
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	// update the newly created post's title
	newTitle := fmt.Sprintf("TestPostsUpdate")
	if post.Title.Raw == newTitle {
		t.Fatalf("New title should be different if we want to test properly")
	}
	post.Title.Raw = newTitle

	// update post
	updatePost, resp, err := wp.Posts.Update(ctx, post.ID, post)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if updatePost == nil {
		t.Errorf("updatePost should not be nil")
	}
	if updatePost.Title.Raw != newTitle {
		t.Errorf("updatePost.Title should be updated to newTitle, %v != %v", updatePost.Title.Raw, newTitle)
	}

	// clea nup
	cleanUpPost(t, updatePost.ID)
}

func TestPostsDelete_NoParams_MoveToTrash(t *testing.T) {
	wp, ctx := initTestClient()

	// create a new post first
	p := factoryPost()
	newPost, resp, _ := wp.Posts.Create(ctx, &p)
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// delete post (move to trash)
	deletedPost, resp, err := wp.Posts.Delete(ctx, newPost.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedPost == nil {
		t.Errorf("updatePost should not be nil")
	}
	if deletedPost.ID != newPost.ID {
		t.Errorf("Deleted post ID should be the same as created post: %v != %v", deletedPost.ID, newPost.ID)
	}

	// clean up
	cleanUpPost(t, newPost.ID)
}
func TestPostsDelete_WithParams_DeletePermanently(t *testing.T) {
	wp, ctx := initTestClient()

	// create a new post first
	p := factoryPost()
	newPost, resp, _ := wp.Posts.Create(ctx, &p)
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// delete post (delete permanently)
	deletedPost, resp, err := wp.Posts.Delete(ctx, newPost.ID, "force=true")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedPost == nil {
		t.Errorf("updatePost should not be nil")
	}
	if deletedPost.ID != newPost.ID {
		t.Errorf("Deleted post ID should be the same as created post: %v != %v", deletedPost.ID, newPost.ID)
	}
}
