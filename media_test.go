package wordpress_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func factoryMediaFileUpload(t *testing.T) *wordpress.MediaUploadOptions {

	// assuming current-working directory `{GO_WORKSPACE_PATH}/src/github.com/robbiet480/go-wordpress`
	path := "./test-data/test-media.jpg"

	// prepare file to upload
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open test media file to upload: %v", err.Error())
	}
	defer file.Close()
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read test media file to upload: %v", err.Error())
	}

	// create / upload media
	media := wordpress.MediaUploadOptions{
		Filename:    "test-media.jpg",
		ContentType: "image/jpeg",
		Data:        fileContents,
	}
	return &media
}

func getAnyOneMedia(t *testing.T, ctx context.Context, wp *wordpress.Client) *wordpress.Media {

	media, resp, _ := wp.Media.List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(media) < 1 {
		t.Fatalf("Should not return empty comments")
	}

	mediaID := media[0].ID

	m, resp, _ := wp.Media.Get(ctx, mediaID, "context=edit")
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if m == nil {
		t.Fatalf("Media should not be nil")
	}
	return m
}

func cleanUpMedia(t *testing.T, ctx context.Context, wp *wordpress.Client, mediaID int) {

	deletedMedia, resp, err := wp.Media.Delete(ctx, mediaID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new media: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedMedia.ID != mediaID {
		t.Errorf("Deleted comment ID should be the same as newly created comment: %v != %v", deletedMedia.ID, mediaID)
	}
}

func TestMediaList(t *testing.T) {

	wp, ctx := initTestClient()

	media, resp, err := wp.Media.List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if media == nil {
		t.Errorf("Should not return nil media")
	}
	if len(media) == 0 {
		t.Errorf("Should not return empty media")
	}
}

func TestMediaGet_Exists(t *testing.T) {
	wp, ctx := initTestClient()

	m := getAnyOneMedia(t, ctx, wp)

	media, resp, err := wp.Media.Get(ctx, m.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if media == nil {
		t.Errorf("Should not return nil media")
	}

}

func TestMediaGet_DoesNotExists(t *testing.T) {
	wp, ctx := initTestClient()

	media, resp, err := wp.Media.Get(ctx, -1, nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp != nil && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %v", resp.Status)
	}

	if media == nil {
		t.Errorf("Should not return nil media")
	}

}

func TestMediaCreate(t *testing.T) {

	wp, ctx := initTestClient()
	media := factoryMediaFileUpload(t)

	newMedia, resp, err := wp.Media.Create(ctx, media)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	if newMedia == nil {
		t.Errorf("Should not return nil newMedia")
	}

	cleanUpMedia(t, ctx, wp, newMedia.ID)
}
