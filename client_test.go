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
	tp := wordpress.BasicAuthTransport{
		Username: USER,
		Password: PASSWORD,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, // needs to be disabled for Lets Encrypt for whatever reason
			DisableKeepAlives: true,
		},
	}
	client := wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: API_BASE_URL,
	}, tp.Client())
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

	tp := wordpress.BasicAuthTransport{
		Username: USER,
		Password: PASSWORD,
		Transport: &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, // needs to be disabled for Lets Encrypt for whatever reason
			DisableKeepAlives: true,
		},
	}

	return wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: API_BASE_URL,
	}, tp.Client()), context.Background()
}
