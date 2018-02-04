package wordpress_test

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func factoryComment(postID int) wordpress.Comment {
	return wordpress.Comment{
		Post:       postID,
		Author:     1,
		Status:     wordpress.CommentStatusApproved,
		AuthorName: "go-wordpress",
		Content: wordpress.RenderedString{
			Raw:      "Test Comment",
			Rendered: "<p>Test Comment</p>",
		},
	}
}

func cleanUpComment(t *testing.T, commentID int) {

	wp, ctx := initTestClient()
	deletedComment, resp, err := wp.Comments.Delete(ctx, commentID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new comment: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedComment.ID != commentID {
		t.Errorf("Deleted comment ID should be the same as newly created comment: %v != %v", deletedComment.ID, commentID)
	}
}

func getAnyOneComment(t *testing.T, ctx context.Context, wp *wordpress.Client) *wordpress.Comment {

	comments, resp, err := wp.Comments.List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(comments) < 1 {
		log.Print(err)
		log.Print(resp)
		t.Fatalf("Should not return empty comments")
	}

	commentID := comments[0].ID

	comment, resp, _ := wp.Comments.Get(ctx, commentID, "context=edit")
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	return comment
}

func TestCommentsList(t *testing.T) {
	wp, ctx := initTestClient()

	comments, resp, err := wp.Comments.List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if comments == nil {
		t.Errorf("Should not return nil comments")
	}
	if len(comments) == 0 {
		t.Errorf("Should not return empty comments")
	}
}

func TestCommentsGet_CommentExists(t *testing.T) {
	wp, ctx := initTestClient()

	c := getAnyOneComment(t, ctx, wp)

	comment, resp, err := wp.Comments.Get(ctx, c.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if comment == nil {
		t.Errorf("Should not return nil comments")
	}
}

func TestCommentsGet_CommentDoesNotExists(t *testing.T) {
	wp, ctx := initTestClient()

	comment, resp, err := wp.Comments.Get(ctx, -1, nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp != nil && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %v", resp.Status)
	}

	if comment == nil {
		t.Errorf("Should not return nil comments")
	}
}

func TestCommentsCreate(t *testing.T) {
	// t.Skipf("[TestCommentsCreate] Skipped: there is an issue with creating comments, server returning empty string")
	wp, ctx := initTestClient()

	p := getAnyOnePost(t, ctx, wp)
	c := factoryComment(p.ID)

	newComment, resp, err := wp.Comments.Create(ctx, &c)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	if newComment == nil {
		t.Errorf("Should not return nil newComment")
	}

	cleanUpComment(t, newComment.ID)
}
