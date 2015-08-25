package wordpress_test

import (
	"github.com/sogko/go-wordpress"
	"os"
	"testing"
)

var USER string = os.Getenv("WP_USER")
var PASSWORD string = os.Getenv("WP_PASSWD")
var API_BASE_URL string = os.Getenv("WP_API_URL")

func TestClientNew(t *testing.T) {
	client := wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: API_BASE_URL,
		Username:   USER,
		Password:   PASSWORD,
	})
	if client == nil {
		t.Fatalf("Client should not be nil")
	}
}

/**
Test helper functions
*/

// initTestClient creates test wordpress client
func initTestClient() *wordpress.Client {
	return wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: API_BASE_URL,
		Username:   USER,
		Password:   PASSWORD,
	})
}
