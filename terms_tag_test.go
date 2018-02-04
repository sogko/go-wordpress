package wordpress_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func factoryTermsTag() *wordpress.Term {
	return &wordpress.Term{
		Name: "TestTermsTagCreate",
		Slug: "TestTermsTagCreate",
	}
}

func cleanUpTermsTag(t *testing.T, id int) {

	wp, ctx := initTestClient()
	deletedTerm, resp, err := wp.Terms.Tag().Delete(ctx, id, nil)
	if err != nil {
		t.Errorf("Failed to clean up new term: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if deletedTerm.ID != id {
		t.Errorf("Deleted term ID should be the same as newly created term: %v != %v", deletedTerm.ID, id)
	}
}

func getAnyOneTermsTag(t *testing.T, ctx context.Context, wp *wordpress.Client) *wordpress.Term {

	terms, resp, _ := wp.Terms.Tag().List(ctx, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(terms) < 1 {
		t.Fatalf("Should not return empty terms")
	}

	id := terms[0].ID

	term, resp, _ := wp.Terms.Tag().Get(ctx, id, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return term
}

func TestTermsTagList(t *testing.T) {
	t.Skipf("Not supported anymore")
	wp, ctx := initTestClient()

	terms, resp, err := wp.Terms.Tag().List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if terms == nil {
		t.Errorf("Should not return nil terms")
	}
}

func TestTermsTagGet(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp, ctx := initTestClient()
	tt := getAnyOneTermsTag(t, ctx, wp)

	term, resp, err := wp.Terms.Tag().Get(ctx, tt.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if term == nil {
		t.Errorf("Should not return nil term")
	}

}

func TestTermsTagCreate(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp, ctx := initTestClient()

	tt := factoryTermsTag()

	term, resp, err := wp.Terms.Tag().Create(ctx, tt)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
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

	wp, ctx := initTestClient()

	tt := factoryTermsTag()

	// create tag
	newTerm, resp, _ := wp.Terms.Tag().Create(ctx, tt)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if newTerm == nil {
		t.Errorf("Should not return nil term")
	}

	// delete tag
	deletedTerm, resp, err := wp.Terms.Tag().Delete(ctx, newTerm.ID, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if deletedTerm == nil {
		t.Errorf("Should not return nil deletedTerm")
	}
}

func TestTermsTagUpdate(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp, ctx := initTestClient()

	tt := factoryTermsTag()

	// create tag
	newTerm, resp, _ := wp.Terms.Tag().Create(ctx, tt)
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if newTerm == nil {
		t.Errorf("Should not return nil term")
	}

	// get tag term
	term, resp, _ := wp.Terms.Tag().Get(ctx, newTerm.ID, nil)
	if resp != nil && resp.StatusCode != http.StatusOK {
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
	updatedTerm, resp, err := wp.Terms.Tag().Update(ctx, newTerm.ID, term)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
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
