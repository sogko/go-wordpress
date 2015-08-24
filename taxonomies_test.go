package wordpress_test

import (
	"net/http"
	"testing"
)

func TestTaxonomiesList(t *testing.T) {
	wp := initTestClient()

	taxonomies, resp, body, err := wp.Taxonomies().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if taxonomies == nil {
		t.Errorf("Should not return nil taxonomies")
	}
	if len(taxonomies) != 2 {
		t.Errorf("Should return two taxonomies")
	}
}

func TestTaxonomiesGet_TaxonomyExists(t *testing.T) {
	wp := initTestClient()

	taxonomy, resp, body, err := wp.Taxonomies().Get("post_tag", nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if taxonomy == nil {
		t.Errorf("Should not return nil taxonomies")
	}
}

func TestTaxonomiesGet_TaxonomyDoesNotExists(t *testing.T) {
	wp := initTestClient()

	taxonomy, resp, body, err := wp.Taxonomies().Get("RANDOM", nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if taxonomy == nil {
		t.Errorf("Should not return nil taxonomies")
	}
}
