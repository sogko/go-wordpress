package wordpress_test

import (
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func cleanUpPageMeta(t *testing.T, page *wordpress.Page, metaId int) {

	// note: Need to pass in `force=true` param in order to delete page meta
	deletedMeta, resp, body, err := page.Meta().Delete(metaId, "force=true")
	if err != nil {
		t.Errorf("Failed to clean up new page: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedMeta.Message != "Deleted meta" {
		t.Errorf("Unexpected response to deleted meta: %v", deletedMeta.Message)
	}

}

func TestPagesMeta_InvalidCall(t *testing.T) {
	// User is not allowed to call create wordpress.Page object manually to retrieve PageMetaCollection
	// A proper API call would inject the right PageMetaCollection, Client and other goodies into a page,
	// allowing user to call page.Meta()
	invalidPage := wordpress.Page{}
	invalidMeta := invalidPage.Meta()
	if invalidMeta != nil {
		t.Error("Expected meta to be nil, %v", invalidMeta)
	}
}

func TestPagesMetaList_NoParams(t *testing.T) {
	wp := initTestClient()

	page := getAnyOnePage(t, wp)

	meta, resp, body, err := page.Meta().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if meta == nil {
		t.Errorf("Should not return nil meta")
	}
}

func TestPagesMetaCreate(t *testing.T) {
	wp := initTestClient()

	// get a page
	page := getAnyOnePage(t, wp)

	// create meta for retrieved page
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, body, err := page.Meta().Create(&m)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	if newMeta == nil {
		t.Errorf("newMeta should not be nil")
	}
	if newMeta.Key != m.Key {
		t.Errorf("newMeta.Key should be the same, %v != %v", newMeta.Key, m.Key)
	}
	if newMeta.Value != m.Value {
		t.Errorf("newMeta.Value should be the same, %v != %v", newMeta.Value, m.Key)
	}

	// clean up
	cleanUpPageMeta(t, page, newMeta.ID)
}

func TestPagesMetaGet(t *testing.T) {
	wp := initTestClient()

	// get a page
	page := getAnyOnePage(t, wp)

	// create meta for retrieved page
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, body, err := page.Meta().Create(&m)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// get meta by id for retrieved page
	metaID := newMeta.ID
	meta, resp, body, err := page.Meta().Get(metaID, nil)
	if err != nil {
		t.Errorf("Failed to get meta: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("meta.Value should be the same, %v != %v", meta.Value, m.Value)
	}
	if meta.ID != metaID {
		t.Errorf("meta.ID should be the same, %v != %v", meta.ID, metaID)
	}
	if newMeta.Key != m.Key {
		t.Errorf("meta.Key should be the same, %v != %v", meta.Key, m.Key)
	}
	if newMeta.Value != m.Value {
		t.Errorf("meta.Value should be the same, %v != %v", meta.Value, m.Value)
	}

	// clean up
	cleanUpPageMeta(t, page, newMeta.ID)
}

func TestPagesMetaGet_Lazy(t *testing.T) {

	wp := initTestClient()

	// get a page so we can have a valid Page ID and we can create a test meta to get
	page := getAnyOnePage(t, wp)
	pageID := page.ID

	// create meta for retrieved page
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, _, _ := page.Meta().Create(&m)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}
	metaID := newMeta.ID

	// Use Pages().Entity(pageID) to retrieve meta in one API call
	lazyMeta, resp, body, err := wp.Pages().Entity(pageID).Meta().Get(metaID, nil)
	if err != nil {
		t.Errorf("Failed to lazy-get meta: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("meta.Value should be the same, %v != %v", lazyMeta.Value, m.Value)
	}
	if lazyMeta.ID != metaID {
		t.Errorf("meta.ID should be the same, %v != %v", lazyMeta.ID, metaID)
	}
	if lazyMeta.Key != m.Key {
		t.Errorf("meta.Key should be the same, %v != %v", lazyMeta.Key, m.Key)
	}
	if lazyMeta.Value != m.Value {
		t.Errorf("meta.Value should be the same, %v != %v", lazyMeta.Value, m.Value)
	}
}

func TestPagesMetaUpdate(t *testing.T) {
	wp := initTestClient()

	// get a page
	page := getAnyOnePage(t, wp)

	// create meta for retrieved page
	m := wordpress.Meta{
		Key:   "testKey",
		Value: "testValue",
	}
	newMeta, resp, body, err := page.Meta().Create(&m)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected 201 Created, got %v", resp.Status)
	}

	// get meta by id for retrieved page
	metaID := newMeta.ID
	meta, resp, body, err := page.Meta().Get(metaID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	// update meta by id for retrieved page
	meta.Value = "newTestValue"
	updatedMeta, resp, body, err := page.Meta().Update(meta.ID, meta)
	if err != nil {
		t.Errorf("Failed to update page meta: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if updatedMeta.ID != meta.ID {
		t.Errorf("updatedMeta.ID should be the same, %v != %v", updatedMeta.ID, meta.ID)
	}
	if updatedMeta.Key != meta.Key {
		t.Errorf("updatedMeta.Key should be the same, %v != %v", updatedMeta.Key, meta.Key)
	}
	if updatedMeta.Value != meta.Value {
		t.Errorf("updatedMeta.Value should be the same, %v != %v", updatedMeta.Value, meta.Value)
	}

	// clean up
	cleanUpPageMeta(t, page, newMeta.ID)
}
