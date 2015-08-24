package wordpress_test

import (
	"github.com/sogko/go-wordpress"
	"net/http"
	"testing"
	"log"
)

func factoryComment(postID int) wordpress.Comment {
	return wordpress.Comment{
		Post: postID,
		Author: 1,
		Status: wordpress.CommentStatusApproved,
		AuthorName: "go-wordpress",
		Content: wordpress.Content{
			Raw: "Test Comment",
			Rendered: "<p>Test Comment</p>",
		},
	}
}

func cleanUpComment(t *testing.T, commentID int) {

	wp := initTestClient()
	deletedComment, resp, body, err := wp.Comments().Delete(commentID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new comment: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedComment.ID != commentID {
		t.Errorf("Deleted comment ID should be the same as newly created comment: %v != %v", deletedComment.ID, commentID)
	}
}

func getAnyOneComment(t *testing.T, wp *wordpress.Client) *wordpress.Comment {

	comments, resp, body, err := wp.Comments().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(comments) < 1 {
		log.Print(err)
		log.Print(body)
		log.Print(resp)
		t.Fatalf("Should not return empty comments")
	}

	commentID := comments[0].ID

	comment, resp, _, _ := wp.Comments().Get(commentID, "context=edit")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	return comment
}

func TestCommentsList(t *testing.T) {
	wp := initTestClient()

	comments, resp, body, err := wp.Comments().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if comments == nil {
		t.Errorf("Should not return nil comments")
	}
	if len(comments) == 0 {
		t.Errorf("Should not return empty comments")
	}
}

func TestCommentsGet_CommentExists(t *testing.T) {
	wp := initTestClient()

	c := getAnyOneComment(t, wp)

	comment, resp, body, err := wp.Comments().Get(c.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if comment == nil {
		t.Errorf("Should not return nil comments")
	}
}

func TestCommentsGet_CommentDoesNotExists(t *testing.T) {
	wp := initTestClient()

	comment, resp, body, err := wp.Comments().Get(-1, nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if comment == nil {
		t.Errorf("Should not return nil comments")
	}
}

func TestCommentsCreate(t *testing.T) {
	t.Skipf("[TestCommentsCreate] Skipped: there is an issue with creating comments, server returning empty string")
	wp := initTestClient()

	p := getAnyOnePost(t, wp)
	c := factoryComment(p.ID)

	newComment, resp, body, err := wp.Comments().Create(&c)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if newComment == nil {
		t.Errorf("Should not return nil newComment")
	}

	cleanUpComment(t, newComment.ID)
}
