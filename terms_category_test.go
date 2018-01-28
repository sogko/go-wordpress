package wordpress_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/robbiet480/go-wordpress"
)

func factoryTermsCategory() *wordpress.Term {
	return &wordpress.Term{
		Name: "TestTermsCategoryCreate4",
		Slug: "TestTermsCategoryCreate4",
	}
}

func cleanUpTermsCategory(t *testing.T, id int) {

	wp := initTestClient()
	deletedTerm, resp, body, err := wp.Terms().Category().Delete(id, nil)
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

func getAnyOneTermsCategory(t *testing.T, wp *wordpress.Client) *wordpress.Term {

	terms, resp, _, _ := wp.Terms().Category().List(nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if len(terms) < 1 {
		t.Fatalf("Should not return empty terms")
	}

	id := terms[0].ID

	term, resp, _, _ := wp.Terms().Category().Get(id, nil)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %v", resp.Status)
	}

	return term
}

func TestTermsCategoryList(t *testing.T) {
	t.Skipf("Not supported anymore")
	wp := initTestClient()

	terms, resp, body, err := wp.Terms().Category().List(nil)
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

func TestTermsCategoryGet(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()
	tt := getAnyOneTermsCategory(t, wp)

	term, resp, body, err := wp.Terms().Category().Get(tt.ID, nil)
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

func TestTermsCategoryCreate_New(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()

	tt := factoryTermsCategory()

	term, resp, body, err := wp.Terms().Category().Create(tt)
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
	cleanUpTermsCategory(t, term.ID)
}

func TestTermsCategoryCreate_Existing(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()

	tt := factoryTermsCategory()

	// add category the first time
	term, resp, body, err := wp.Terms().Category().Create(tt)
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

	// add the same category the second time
	duplicateTerm, resp, body, err := wp.Terms().Category().Create(tt)
	if err == nil {
		t.Errorf("Should return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected 500 Internal Server Erro, got %v", resp.Status)
	}
	if duplicateTerm == nil {
		t.Errorf("Should not return nil duplicateTerm")
	}

	// unmarshall error response
	// We expect server to return "term_exists" error code
	serverErrors, err := wordpress.UnmarshallServerError(body)
	if err != nil {
		cleanUpTermsCategory(t, term.ID)
		log.Println(string(body))
		t.Fatalf("Unexpected error response from server, unable to unmarshall message %v", err.Error())
	}
	if len(serverErrors) != 1 {
		t.Error("Expected one error", len(serverErrors))
	}
	if serverErrors[0].Code != "term_exists" {
		t.Errorf("Unexpected err.code, %v != term_exists", serverErrors[0].Code)
	}

	// clean up
	cleanUpTermsCategory(t, term.ID)

}

func TestTermsCategoryDelete(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()

	tt := factoryTermsCategory()

	// create category
	newTerm, resp, _, _ := wp.Terms().Category().Create(tt)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if newTerm == nil {
		t.Errorf("Should not return nil term")
	}

	// delete category
	deletedTerm, resp, body, err := wp.Terms().Category().Delete(newTerm.ID, nil)
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

func TestTermsCategoryUpdate(t *testing.T) {
	t.Skipf("Not supported anymore")

	wp := initTestClient()

	tt := factoryTermsCategory()

	// create category
	newTerm, resp, _, _ := wp.Terms().Category().Create(tt)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if newTerm == nil {
		t.Errorf("Should not return nil term")
	}

	// get category term
	term, resp, _, _ := wp.Terms().Category().Get(newTerm.ID, nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if term == nil {
		t.Errorf("Should not return nil term")
	}

	// modify term description
	newTermDescription := "TestTermsCategoryUpdate"
	if term.Description == newTermDescription {
		t.Errorf("Warning: Data must be different for proper test, %v === %v", term.Description, newTermDescription)
	}
	term.Description = newTermDescription

	// update
	updatedTerm, resp, body, err := wp.Terms().Category().Update(newTerm.ID, term)
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
	cleanUpTermsCategory(t, newTerm.ID)
}
