package wordpress_test

import (
	"fmt"
	"github.com/sogko/go-wordpress"
	"net/http"
	"testing"
	"time"
)

func getLatestRevisionForPage(t *testing.T, page *wordpress.Page) *wordpress.Revision {

	revisions, resp, _, _ := page.Revisions().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(revisions) < 1 {
		t.Fatalf("Should not return empty revisions")
	}
	// get latest revision
	revisionID := revisions[0].ID
	revision, resp, _, _ := page.Revisions().Get(revisionID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return revision
}

func TestPagesRevisions_InvalidCall(t *testing.T) {
	// User is not allowed to call create wordpress.Page object manually to retrieve PageMetaCollection
	// A proper API call would inject the right PageMetaCollection, Client and other goodies into a page,
	// allowing user to call page.Revisions()
	invalidPage := wordpress.Page{}
	invalidRevisions := invalidPage.Revisions()
	if invalidRevisions != nil {
		t.Error("Expected revisions to be nil, %v", invalidRevisions)
	}
}

func TestPagesRevisionsList(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)

	revisions, resp, body, err := page.Revisions().List(nil)
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

func TestPagesRevisionsList_Lazy(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)
	pageID := page.ID

	// Use Pages().Entity(pageID) to retrieve revisions in one API call
	lazyRevisions, resp, body, err := wp.Pages().Entity(pageID).Revisions().List(nil)
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

func TestPagesRevisionsGet(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)
	r := getLatestRevisionForPage(t, page)

	revisionID := r.ID

	revision, resp, body, err := page.Revisions().Get(revisionID, nil)
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

func TestPagesRevisionsGet_Lazy(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)
	r := getLatestRevisionForPage(t, page)

	pageID := page.ID
	revisionID := r.ID

	// Use Pages().Entity(pageID) to retrieve revisions in one API call
	lazyRevision, resp, body, err := wp.Pages().Entity(pageID).Revisions().Get(revisionID, nil)
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

func TestPagesRevisionsDelete_Lazy(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)

	// Edit page to create a new revision
	// Note: wordpress would only create a new revision if there is an actual change in
	// content
	now := time.Now()
	originalTitle := page.Title.Raw
	page.Title.Raw = fmt.Sprintf("%v", now.Format("20060102150405"))
	if originalTitle == page.Title.Raw {
		t.Fatalf("Flawed test, ensure that page content is modified before an update")
	}
	updatedPage, resp, _, _ := wp.Pages().Update(page.ID, page)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	r := getLatestRevisionForPage(t, updatedPage)
	pageID := updatedPage.ID
	revisionID := r.ID

	// Use Pages().Entity(pageID) to delete revisions in one API call
	// Note that deleting a revision does NOT reverse the changes made in the revision
	response, resp, body, err := wp.Pages().Entity(pageID).Revisions().Delete(revisionID, nil)
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
