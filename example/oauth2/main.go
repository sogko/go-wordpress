package main

import (
	"context"
	"log"

	"github.com/robbiet480/go-wordpress"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "JWT_TOKEN"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: "http://192.168.99.100:32777/wp-json/",
	}, tc)

	// get the currently authenticated users details
	authenticatedUser, _, err := client.Users.Me(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Authenticated user %+v", authenticatedUser)
}
