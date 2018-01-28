package wordpress_test

import (
	"os"
	"testing"

	"github.com/robbiet480/go-wordpress"
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

	if API_BASE_URL == "" {
		panic("Please set your environment before running the tests")
	}

	return wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: API_BASE_URL,
		Username:   USER,
		Password:   PASSWORD,
	})
}
