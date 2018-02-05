package main

import (
	"context"
	"log"

	"github.com/robbiet480/go-wordpress"
)

func main() {

	tp := wordpress.BasicAuthTransport{
		Username: "username",
		Password: "password",
	}

	// create wp-api client
	client := wordpress.NewClient(&wordpress.Options{
		BaseAPIURL: "http://192.168.99.100:32777/wp-json/",
	}, tp.Client())

	ctx := context.Background()

	// get the currently authenticated users details
	authenticatedUser, _, err := client.Users.Me(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Authenticated user %+v", authenticatedUser)
}
