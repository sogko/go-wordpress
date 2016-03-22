package wordpress_test

import (
	"github.com/sogko/go-wordpress"
	"net/http"
	"testing"
)

func factoryTermsTag() *wordpress.Term {
	return &wordpress.Term{
		Name: "TestTermsTagCreate",
		Slug: "TestTermsTagCreate",
	}
}

func cleanUpTermsTag(t *testing.T, id int) {

	wp := initTestClient()
	deletedTerm, resp, body, err := wp.Terms().Tag().Delete(id, nil)
	if err != nil {
		t.Errorf("Failed to clean up new term: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedTerm.ID != id {
		t.Errorf("Deleted term ID should be the same as newly created term: %v != %v", deletedTerm.ID, id)
	}
}

func getAnyOneTermsTag(t *testing.T, wp *wordpress.Client) *wordpress.Term {

	terms, resp, _, _ := wp.Terms().Tag().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(terms) < 1 {
		t.Fatalf("Should not return empty terms")
	}

	id := terms[0].ID

	term, resp, _, _ := wp.Terms().Tag().Get(id, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return term
}

func TestTermsTagList(t *testing.T) {
	t.Skipf("Not supported anymore")
	wp := initTestClient()

	terms, resp, body, err := wp.Terms().Tag().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if terms == nil {
		t.Errorf("Should not return nil terms")
	}
}

func TestTermsTagGet(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()
	tt := getAnyOneTermsTag(t, wp)

	term, resp, body, err := wp.Terms().Tag().Get(tt.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if term == nil {
		t.Errorf("Should not return nil term")
	}

}

func TestTermsTagCreate(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()

	tt := factoryTermsTag()

	term, resp, body, err := wp.Terms().Tag().Create(tt)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if term == nil {
		t.Errorf("Should not return nil term")
	}

	// clean up
	cleanUpTermsTag(t, term.ID)
}

func TestTermsTagDelete(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()

	tt := factoryTermsTag()

	// create tag
	newTerm, resp, _, _ := wp.Terms().Tag().Create(tt)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if newTerm == nil {
		t.Errorf("Should not return nil term")
	}

	// delete tag
	deletedTerm, resp, body, err := wp.Terms().Tag().Delete(newTerm.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedTerm == nil {
		t.Errorf("Should not return nil deletedTerm")
	}
}

func TestTermsTagUpdate(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()

	tt := factoryTermsTag()

	// create tag
	newTerm, resp, _, _ := wp.Terms().Tag().Create(tt)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if newTerm == nil {
		t.Errorf("Should not return nil term")
	}

	// get tag term
	term, resp, _, _ := wp.Terms().Tag().Get(newTerm.ID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if term == nil {
		t.Errorf("Should not return nil term")
	}

	// modify term description
	newTermDescription := "TestTermsTagUpdate"
	if term.Description == newTermDescription {
		t.Errorf("Warning: Data must be different for proper test, %v === %v", term.Description, newTermDescription)
	}
	term.Description = newTermDescription

	// update
	updatedTerm, resp, body, err := wp.Terms().Tag().Update(newTerm.ID, term)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if updatedTerm == nil {
		t.Errorf("Should not return nil updatedTerm")
	}
	if updatedTerm.Description != newTermDescription {
		t.Errorf("Expected term to have updated description")
	}

	// clean up
	cleanUpTermsTag(t, newTerm.ID)
}
