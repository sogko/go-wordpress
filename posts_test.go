package wordpress_test

import (
	"fmt"
	"github.com/sogko/go-wordpress"
	"log"
	"net/http"
	"testing"
)

func factoryPost() wordpress.Post {
	return wordpress.Post{
		Title: wordpress.Title{
			Raw: "TestPostsCreate",
		},
		Content: wordpress.Content{
			Raw: "<h1>HEADER</h1><p>Paragraph</p>",
		},
		Excerpt: wordpress.Excerpt{
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

	wp := initTestClient()
	deletedPost, resp, body, err := wp.Posts().Delete(postID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new post: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedPost.ID != postID {
		t.Errorf("Deleted post ID should be the same as newly created post: %v != %v", deletedPost.ID, postID)
	}
}

func getAnyOnePost(t *testing.T, wp *wordpress.Client) *wordpress.Post {

	posts, resp, body, err := wp.Posts().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(posts) < 1 {
		log.Print(err)
		log.Print(body)
		log.Print(resp)
		t.Fatalf("Should not return empty posts")
	}

	postID := posts[0].ID

	post, resp, _, _ := wp.Posts().Get(postID, "context=edit")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	return post
}

func TestPostsList_NoParams(t *testing.T) {
	wp := initTestClient()

	posts, resp, body, err := wp.Posts().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if posts == nil {
		t.Errorf("Should not return nil posts")
	}
	if len(posts) == 0 {
		t.Errorf("Should not return empty posts")
	}
}
func TestPostsList_WithParamsString(t *testing.T) {
	wp := initTestClient()

	// assumes that API user authenticated with `edit_posts`
	posts, resp, body, err := wp.Posts().List("filter[post_status]=draft")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if len(posts) != 0 {
		t.Errorf("Should return zero draft posts, returned %v", len(posts))
	}
	posts, resp, body, err = wp.Posts().List("filter[post_status]=publish")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if len(posts) == 0 {
		t.Errorf("Should return at least one published posts")
	}
}

func TestPostsGet_PostExists(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)
	postID := post.ID

	post, resp, body, err := wp.Posts().Get(postID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if post.ID != postID {
		t.Errorf("Returned post should have the same ID as specified in Get(), %v != %v", post.ID, postID)
	}
}

func TestPostsGet_PostDoesNotExists(t *testing.T) {
	wp := initTestClient()

	postID := -1

	_, resp, body, err := wp.Posts().Get(postID, nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 400 NotFound, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
}

func TestPostsGet_Lazy(t *testing.T) {
	wp := initTestClient()

	post := getAnyOnePost(t, wp)
	postID := post.ID

	//The proper way to get lazy-fetch posts. Posts().Entity() won't make any HTTP request
	lazyPost := wp.Posts().Entity(postID)
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
	post, resp, body, err := lazyPost.Populate(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if post.ID != postID {
		t.Errorf("Returned post should have the same ID as specified in Get(), %v != %v", post.ID, postID)
	}
	if post.GUID.Rendered == "" {
		t.Errorf("post should have populated GUID field, %v", lazyPost.GUID.Rendered)
	}
}

func TestPostsCreate(t *testing.T) {
	wp := initTestClient()

	p := factoryPost()
	newPost, resp, body, err := wp.Posts().Create(&p)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
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
	wp := initTestClient()

	// create a new post first
	p := factoryPost()
	newPost, resp, _, _ := wp.Posts().Create(&p)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %v", resp.Status)
	}

	// get the post in `edit` context
	post, resp, _, _ := wp.Posts().Get(newPost.ID, "context=edit")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	// update the newly created post's title
	newTitle := fmt.Sprintf("TestPostsUpdate")
	if post.Title.Raw == newTitle {
		t.Fatalf("New title should be different if we want to test properly")
	}
	post.Title.Raw = newTitle

	// update post
	updatePost, resp, body, err := wp.Posts().Update(post.ID, post)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
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
	wp := initTestClient()

	// create a new post first
	p := factoryPost()
	newPost, resp, _, _ := wp.Posts().Create(&p)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// delete post (move to trash)
	deletedPost, resp, body, err := wp.Posts().Delete(newPost.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
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
	wp := initTestClient()

	// create a new post first
	p := factoryPost()
	newPost, resp, _, _ := wp.Posts().Create(&p)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// delete post (delete permanently)
	deletedPost, resp, body, err := wp.Posts().Delete(newPost.ID, "force=true")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if deletedPost == nil {
		t.Errorf("updatePost should not be nil")
	}
	if deletedPost.ID != newPost.ID {
		t.Errorf("Deleted post ID should be the same as created post: %v != %v", deletedPost.ID, newPost.ID)
	}
}
