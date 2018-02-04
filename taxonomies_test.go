package wordpress_test

import (
	"net/http"
	"testing"
)

func TestTaxonomiesList(t *testing.T) {
	wp, ctx := initTestClient()

	taxonomies, resp, err := wp.Taxonomies.List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if taxonomies == nil {
		t.Errorf("Should not return nil taxonomies")
	}
	if len(taxonomies) != 2 {
		t.Errorf("Should return two taxonomies")
	}
}

func TestTaxonomiesGet_TaxonomyExists(t *testing.T) {
	wp, ctx := initTestClient()

	taxonomy, resp, err := wp.Taxonomies.Get(ctx, "post_tag", nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if taxonomy == nil {
		t.Errorf("Should not return nil taxonomies")
	}
}

func TestTaxonomiesGet_TaxonomyDoesNotExists(t *testing.T) {
	wp, ctx := initTestClient()

	taxonomy, resp, err := wp.Taxonomies.Get(ctx, "RANDOM", nil)
	if err == nil {
		t.Errorf("Should return error")
	}
	if resp != nil && resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected 404 Not Found, got %v", resp.Status)
	}

	if taxonomy == nil {
		t.Errorf("Should not return nil taxonomies")
	}
}
