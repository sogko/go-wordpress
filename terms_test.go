package wordpress_test

import (
	"net/http"
	"testing"
)

func TestTermsList(t *testing.T) {
	t.Skipf("Not supported anymore")
	wp, ctx := initTestClient()

	terms, resp, err := wp.Terms.List(ctx, "tag", nil)
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
