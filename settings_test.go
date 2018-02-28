package wordpress_test

import (
	"net/http"
	"testing"
)

func TestSettingsList(t *testing.T) {
	client, ctx := initTestClient()

	settings, resp, err := client.Settings.List(ctx)
	if err != nil {
		t.Errorf("Should not return error: %v", err.Error())
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 StatusOK, got %v", resp.Status)
	}
	if settings == nil {
		t.Errorf("Should not return nil settings")
	}
}
