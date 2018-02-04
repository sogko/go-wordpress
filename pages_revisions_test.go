package wordpress_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/robbiet480/go-wordpress"
)

func getLatestRevisionForPage(t *testing.T, ctx context.Context, page *wordpress.Page) *wordpress.Revision {

	revisions, resp, _ := page.Revisions().List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(revisions) < 1 {
		t.Fatalf("Should not return empty revisions")
	}
	// get latest revision
	revisionID := revisions[0].ID
	revision, resp, _ := page.Revisions().Get(ctx, revisionID, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return revision
}

func TestPagesRevisions_InvalidCall(t *testing.T) {
	// User is not allowed to call create wordpress.Page object manually to retrieve PageMetaService
	// A proper API call would inject the right PageMetaService, Client and other goodies into a page,
	// allowing user to call page.Revisions
	invalidPage := wordpress.Page{}
	invalidRevisions := invalidPage.Revisions()
	if invalidRevisions != nil {
		t.Errorf("Expected revisions to be nil, %v", invalidRevisions)
	}
}

func TestPagesRevisionsList(t *testing.T) {
	wp, ctx := initTestClient()

	page := getAnyOnePage(t, ctx, wp)

	revisions, resp, err := page.Revisions().List(ctx, nil)
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

func TestPagesRevisionsList_Lazy(t *testing.T) {
	wp, ctx := initTestClient()

	page := getAnyOnePage(t, ctx, wp)
	pageID := page.ID

	// Use Pages.Entity(pageID) to retrieve revisions in one API call
	lazyRevisions, resp, err := wp.Pages.Entity(pageID).Revisions().List(ctx, nil)
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

func TestPagesRevisionsGet(t *testing.T) {
	wp, ctx := initTestClient()

	page := getAnyOnePage(t, ctx, wp)
	r := getLatestRevisionForPage(t, ctx, page)

	revisionID := r.ID

	revision, resp, err := page.Revisions().Get(ctx, revisionID, nil)
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

func TestPagesRevisionsGet_Lazy(t *testing.T) {
	wp, ctx := initTestClient()

	page := getAnyOnePage(t, ctx, wp)
	r := getLatestRevisionForPage(t, ctx, page)

	pageID := page.ID
	revisionID := r.ID

	// Use Pages.Entity(pageID) to retrieve revisions in one API call
	lazyRevision, resp, err := wp.Pages.Entity(pageID).Revisions().Get(ctx, revisionID, nil)
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

func TestPagesRevisionsDelete_Lazy(t *testing.T) {
	wp, ctx := initTestClient()

	page := getAnyOnePage(t, ctx, wp)

	// Edit page to create a new revision
	// Note: wordpress would only create a new revision if there is an actual change in
	// content
	now := time.Now()
	originalTitle := page.Title.Raw
	page.Title.Raw = fmt.Sprintf("%v", now.Format("20060102150405"))
	if originalTitle == page.Title.Raw {
		t.Fatalf("Flawed test, ensure that page content is modified before an update")
	}
	updatedPage, resp, _ := wp.Pages.Update(ctx, page.ID, page)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	r := getLatestRevisionForPage(t, ctx, updatedPage)
	pageID := updatedPage.ID
	revisionID := r.ID

	// Use Pages.Entity(pageID) to delete revisions in one API call
	// Note that deleting a revision does NOT reverse the changes made in the revision
	response, resp, err := wp.Pages.Entity(pageID).Revisions().Delete(ctx, revisionID, nil)
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
