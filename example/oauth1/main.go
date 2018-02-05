package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dghubble/oauth1"
	"github.com/robbiet480/go-wordpress"
)

var config oauth1.Config

// main performs the WordPress OAuth1 user flow from the command line
func main() {
	config = oauth1.Config{
		ConsumerKey:    "CONSUMER_KEY",
		ConsumerSecret: "CONSUMER_SECRET",
		CallbackURL:    "http://localhost:8080/callback",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "http://192.168.99.100:32777/oauth1/request",
			AuthorizeURL:    "http://192.168.99.100:32777/oauth1/authorize",
			AccessTokenURL:  "http://192.168.99.100:32777/oauth1/access",
		},
	}

	requestToken, requestSecret, err := login()
	if err != nil {
		log.Fatalf("Request Token Phase: %s", err.Error())
	}
	accessToken, err := receiveVerifier(requestToken, requestSecret)
	if err != nil {
		log.Fatalf("Access Token Phase: %s", err.Error())
	}

	log.Println("Consumer was granted an access token to act on behalf of a user.")
	log.Printf("token: %s\nsecret: %s\n", accessToken.Token, accessToken.TokenSecret)

	ctx := context.Background()

	httpClient := config.Client(ctx, accessToken)

	// create wp-api client
	client, _ := wordpress.NewClient("http://192.168.99.100:32777/wp-json/", httpClient)

	// get the currently authenticated users details
	authenticatedUser, _, err := client.Users.Me(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Authenticated user %+v", authenticatedUser)
}

func login() (requestToken, requestSecret string, err error) {
	requestToken, requestSecret, err = config.RequestToken()
	if err != nil {
		return "", "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", "", err
	}
	fmt.Printf("Open this URL in your browser:\n%s\n", authorizationURL.String())
	return requestToken, requestSecret, err
}

func receiveVerifier(requestToken, requestSecret string) (*oauth1.Token, error) {
	fmt.Printf("Choose whether to grant the application access.\nPaste " +
		"the oauth_verifier parameter from the address bar: ")
	var verifier string
	_, err := fmt.Scanf("%s", &verifier)
	accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, verifier)
	if err != nil {
		return nil, err
	}
	return oauth1.NewToken(accessToken, accessSecret), err
}
