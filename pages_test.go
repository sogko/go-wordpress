package wordpress_test

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func factoryPage() wordpress.Page {
	return wordpress.Page{
		Title: wordpress.Title{
			Raw: "TestPagesCreate",
		},
		Content: wordpress.Content{
			Raw: "<h1>HEADER</h1><p>Paragraph</p>",
		},
		Excerpt: wordpress.Excerpt{
			Raw: "<h1>HEADER</h1><p>Paragraph</p>",
		},
		Type:   wordpress.PostTypePage,
		Status: wordpress.PostStatusDraft,
		Slug:   "test-pages-create",
		Author: 1,
	}
}

func cleanUpPage(t *testing.T, pageID int) {

	wp := initTestClient()
	deletedPage, resp, body, err := wp.Pages().Delete(pageID, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new page: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedPage.ID != pageID {
		t.Errorf("Deleted page ID should be the same as newly created page: %v != %v", deletedPage.ID, pageID)
	}
}

func getAnyOnePage(t *testing.T, wp *wordpress.Client) *wordpress.Page {

	pages, resp, body, err := wp.Pages().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(pages) < 1 {
		log.Print(err)
		log.Print(body)
		log.Print(resp)
		t.Fatalf("Should not return empty pages")
	}

	pageID := pages[0].ID

	page, resp, _, _ := wp.Pages().Get(pageID, "context=edit")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	return page
}

func TestPagesList_NoParams(t *testing.T) {
	wp := initTestClient()

	pages, resp, body, err := wp.Pages().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if pages == nil {
		t.Errorf("Should not return nil pages")
	}
	if len(pages) == 0 {
		t.Errorf("Should not return empty pages")
	}
}
func TestPagesList_WithParamsString(t *testing.T) {
	wp := initTestClient()

	// assumes that API user authenticated with `edit_pages`
	pages, resp, body, err := wp.Pages().List("filter[post_status]=draft")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if len(pages) != 0 {
		t.Errorf("Should return zero draft pages, returned %v", len(pages))
	}
	pages, resp, body, err = wp.Pages().List("filter[post_status]=publish")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if len(pages) == 0 {
		t.Errorf("Should return at least one published pages")
	}
}

func TestPagesGet_PageExists(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)
	pageID := page.ID

	page, resp, body, err := wp.Pages().Get(pageID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if page.ID != pageID {
		t.Errorf("Returned page should have the same ID as specified in Get(), %v != %v", page.ID, pageID)
	}
}
func TestPagesGet_PageDoesNotExists(t *testing.T) {
	wp := initTestClient()

	pageID := -1

	_, resp, body, err := wp.Pages().Get(pageID, nil)
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
func TestPagesGet_Lazy(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)
	pageID := page.ID

	//The proper way to get lazy-fetch pages. Pages().Entity() won't make any HTTP request
	lazyPage := wp.Pages().Entity(pageID)
	if lazyPage == nil {
		t.Errorf("lazyPage should not be nil")
	}
	if lazyPage.ID != pageID {
		t.Errorf("lazyPage should have specified ID, %v != %v", lazyPage.ID, pageID)
	}
	if lazyPage.GUID.Rendered != "" {
		t.Errorf("lazyPage should not have populated GUID field, %v", lazyPage.GUID.Rendered)
	}

	// populate Page Entity
	page, resp, body, err := lazyPage.Populate(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if page.ID != pageID {
		t.Errorf("Returned page should have the same ID as specified in Get(), %v != %v", page.ID, pageID)
	}
	if page.GUID.Rendered == "" {
		t.Errorf("page should have populated GUID field, %v", lazyPage.GUID.Rendered)
	}
}

func TestPagesCreate(t *testing.T) {
	wp := initTestClient()

	p := factoryPage()
	newPage, resp, body, err := wp.Pages().Create(&p)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if newPage == nil {
		t.Errorf("newPage should not be nil")
	}
	if newPage.ID < 1 {
		t.Errorf("newPage.ID should not be invalid")
	}
	if newPage.Title.Raw != p.Title.Raw {
		t.Errorf("newPage.Title should be the same, %v != %v", newPage.Title.Raw, p.Title.Raw)
	}
	if newPage.Status != p.Status {
		t.Errorf("newPage.Status should be the same, %v != %v", newPage.Status, p.Status)
	}
	if newPage.Slug != p.Slug {
		t.Errorf("newPage.Slug should be the same, %v != %v", newPage.Slug, p.Slug)
	}

	// clean up
	cleanUpPage(t, newPage.ID)
}

func TestPagesUpdate(t *testing.T) {
	wp := initTestClient()

	// create a new page first
	p := factoryPage()
	newPage, resp, _, _ := wp.Pages().Create(&p)
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %v", resp.Status)
	}

	// get the page in `edit` context
	page, resp, _, _ := wp.Pages().Get(newPage.ID, "context=edit")
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	// update the newly created page's title
	newTitle := fmt.Sprintf("TestPagesUpdate")
	if page.Title.Raw == newTitle {
		t.Fatalf("New title should be different if we want to test properly")
	}
	page.Title.Raw = newTitle

	// update page
	updatePage, resp, body, err := wp.Pages().Update(page.ID, page)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if updatePage == nil {
		t.Errorf("updatePage should not be nil")
	}
	if updatePage.Title.Raw != newTitle {
		t.Errorf("updatePage.Title should be updated to newTitle, %v != %v", updatePage.Title.Raw, newTitle)
	}

	// clea nup
	cleanUpPage(t, updatePage.ID)
}

func TestPagesDelete_NoParams_MoveToTrash(t *testing.T) {
	wp := initTestClient()

	// create a new page first
	p := factoryPage()
	newPage, resp, _, _ := wp.Pages().Create(&p)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// delete page (move to trash)
	deletedPage, resp, body, err := wp.Pages().Delete(newPage.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if deletedPage == nil {
		t.Errorf("updatePage should not be nil")
	}
	if deletedPage.ID != newPage.ID {
		t.Errorf("Deleted page ID should be the same as created page: %v != %v", deletedPage.ID, newPage.ID)
	}

	// clean up
	cleanUpPage(t, newPage.ID)
}
func TestPagesDelete_WithParams_DeletePermanently(t *testing.T) {
	wp := initTestClient()

	// create a new page first
	p := factoryPage()
	newPage, resp, _, _ := wp.Pages().Create(&p)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// delete page (delete permanently)
	deletedPage, resp, body, err := wp.Pages().Delete(newPage.ID, "force=true")
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if deletedPage == nil {
		t.Errorf("updatePage should not be nil")
	}
	if deletedPage.ID != newPage.ID {
		t.Errorf("Deleted page ID should be the same as created page: %v != %v", deletedPage.ID, newPage.ID)
	}
}
