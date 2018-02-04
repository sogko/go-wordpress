package wordpress_test

import (
	"context"
	"crypto/tls"
	"net/http"
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
	}, nil)
	if client == nil {
		t.Fatalf("Client should not be nil")
	}
}

/**
Test helper functions
*/

// initTestClient creates test wordpress client
func initTestClient() (*wordpress.Client, context.Context) {

	if API_BASE_URL == "" {
		panic("Please set your environment before running the tests")
	}

	httpClient := &http.Client{
		Jar: nil,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, // needs to be disabled for Lets Encrypt for whatever reason
			DisableKeepAlives: true,
		},
	}

	httpClient.CheckRedirect = func(r *http.Request, via []*http.Request) error {
		r.SetBasicAuth(USER, PASSWORD)
		return nil
	}

	return wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: API_BASE_URL,
		Username:   USER,
		Password:   PASSWORD,
	}, httpClient), context.Background()
}
