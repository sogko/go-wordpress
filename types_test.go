package wordpress_test

import (
	"net/http"
	"testing"
)

func TestTypesList(t *testing.T) {
	wp, ctx := initTestClient()

	types, resp, err := wp.Types.List(ctx, nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if types == nil {
		t.Errorf("Should not return nil types")
	}
}

func TestTypesGet(t *testing.T) {
	wp, ctx := initTestClient()

	wpType, resp, err := wp.Types.Get(ctx, "post", nil)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", resp.Status)
	}

	if wpType == nil {
		t.Errorf("Should not return nil type")
	}
}
