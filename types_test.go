package wordpress_test

import (
	"net/http"
	"testing"
)


func TestTypesList(t *testing.T) {
	wp := initTestClient()

	types, resp, body, err := wp.Types().List(nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if types == nil {
		t.Errorf("Should not return nil types")
	}
}

func TestTypesGet(t *testing.T) {
	wp := initTestClient()

	wpType, resp, body, err := wp.Types().Get("post", nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}
	if body == nil {
		t.Errorf("Should not return nil body")
	}
	if wpType == nil {
		t.Errorf("Should not return nil type")
	}
}
