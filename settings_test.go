package wordpress_test

import (
	"net/http"
	"testing"
)

func TestSettingsList(t *testing.T) {
	client := initTestClient()

	settings, resp, body, err := client.Settings().List()
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if body == nil {
		t.Errorf("body should not be nil")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if settings == nil {
		t.Errorf("Should not return nil settings")
	}
}
