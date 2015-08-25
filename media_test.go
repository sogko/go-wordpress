package wordpress_test
import (
	"testing"
	"net/http"
	"github.com/sogko/go-wordpress"
	"os"
	"io/ioutil"
)

func factoryMediaFileUpload(t *testing.T) *wordpress.MediaUploadOptions {

	// assuming current-working directory `{GO_WORKSPACE_PATH}/src/github.com/sogko/go-wordpress`
	path := "./test-data/test-media.jpg"

	// prepare file to upload
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		t.Fatalf("Failed to open test media file to upload: %v", err.Error())
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read test media file to upload: %v", err.Error())
	}

	// create / upload media
	media := wordpress.MediaUploadOptions{
		Filename: "test-media.jpg",
		ContentType: "image/jpeg",
		Data: fileContents,
	}
	return &media
}
func getAnyOneMedia(t *testing.T, wp *wordpress.Client) *wordpress.Media {

	media, resp, _, _ := wp.Media().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(media) < 1 {
		t.Fatalf("Should not return empty comments")
	}

	mediaID := media[0].ID

	m, resp, _, _ := wp.Media().Get(mediaID, "context=edit")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if m == nil {
		t.Fatalf("Media should not be nil")
	}
	return m
}

func cleanUpMedia(t *testing.T, wp *wordpress.Client, mediaID int) {

	deletedMedia, resp, body, err := wp.Media().Delete(mediaID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new media: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedMedia.ID != mediaID {
		t.Errorf("Deleted comment ID should be the same as newly created comment: %v != %v", deletedMedia.ID, mediaID)
	}
}
func TestMediaList(t *testing.T) {

	wp := initTestClient()

	media, resp, body, err := wp.Media().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if media == nil {
		t.Errorf("Should not return nil media")
	}
	if len(media) == 0 {
		t.Errorf("Should not return empty media")
	}
}

func TestMediaGet_Exists(t *testing.T) {
	wp := initTestClient()

	m := getAnyOneMedia(t, wp)

	media, resp, body, err := wp.Media().Get(m.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if media == nil {
		t.Errorf("Should not return nil media")
	}

}

func TestMediaGet_DoesNotExists(t *testing.T) {
	wp := initTestClient()

	media, resp, body, err := wp.Media().Get(-1, nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if media == nil {
		t.Errorf("Should not return nil media")
	}

}

func TestMediaCreate(t *testing.T) {

	wp := initTestClient()
	media := factoryMediaFileUpload(t)

	newMedia, resp, body, err := wp.Media().Create(media)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if newMedia == nil {
		t.Errorf("Should not return nil newMedia")
	}

	cleanUpMedia(t, wp, newMedia.ID)
}
