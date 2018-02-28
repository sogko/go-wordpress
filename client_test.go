package wordpress_test

import (
	"context"
	"errors"
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
	}
	client, clientErr := wordpress.NewClient(API_BASE_URL, tp.Client())
	if client == nil {
		t.Fatalf("Client should not be nil")
	}

	if clientErr != nil {
		t.Fatal("Error parsing URL")
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
	}

	client, clientErr := wordpress.NewClient(API_BASE_URL, tp.Client())

	if client == nil {
		panic(errors.New("client should not be nil"))
	}

	if clientErr != nil {
		panic(errors.New("error parsing url"))
	}

	return client, context.Background()
}
