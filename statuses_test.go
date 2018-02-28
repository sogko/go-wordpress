package wordpress_test

import (
	"net/http"
	"testing"
)

func TestStatusesList(t *testing.T) {
	wp, ctx := initTestClient()

	statuses, resp, err := wp.Statuses.List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if statuses == nil {
		t.Errorf("Should not return nil statuses")
	}
}

func TestStatusesGet(t *testing.T) {
	wp, ctx := initTestClient()

	status, resp, err := wp.Statuses.Get(ctx, "publish", nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if status == nil {
		t.Errorf("Should not return nil status")
	}
}
