package wordpress_test

import (
	"net/http"
	"testing"
)


func TestStatusesList(t *testing.T) {
	wp := initTestClient()

	statuses, resp, body, err := wp.Statuses().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if statuses == nil {
		t.Errorf("Should not return nil statuses")
	}
}

func TestStatusesGet(t *testing.T) {
	wp := initTestClient()

	status, resp, body, err := wp.Statuses().Get("publish", nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if status == nil {
		t.Errorf("Should not return nil status")
	}
}
