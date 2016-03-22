package wordpress_test

import (
	"net/http"
	"testing"
)

func TestTermsList(t *testing.T) {
	t.Skipf("Not supported anymore")
	wp := initTestClient()

	terms, resp, body, err := wp.Terms().List("tag", nil)
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
