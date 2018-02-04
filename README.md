# go-wp-api
[![GoDoc](https://godoc.org/github.com/robbiet480/go-wordpress?status.svg)](https://godoc.org/github.com/robbiet480/go-wordpress)

A Go client library for the [Wordpress REST API](https://developer.wordpress.org/rest-api/)


## Installation

```bash
go get github.com/robbiet480/go-wordpress

```

## Usage

### Quick example
```go
package main

import (
	"github.com/robbiet480/go-wordpress"
	"net/http"
)

func main() {

  // create wp-api client
  client := wordpress.NewClient(&wordpress.Options{
    BaseAPIURL: API_BASE_URL, // example: `http://192.168.99.100:32777/wp-json/`
    Username:   USER,
    Password:   PASSWORD,
  }, nil)

  ctx := context.Background()

  // for eg, to get current user (GET /users/me)
  currentUser, resp, _ := client.Users.Me(ctx)
  if resp != nil && resp.StatusCode != http.StatusOK {
    // handle error
  }

  // Or you can use your own structs (for custom endpoints, for example)
  // Below is the equivalent of `client.Posts.Get(100, nil)`
  var obj MyCustomPostStruct
  resp, err := client.Get(ctx, "/posts/100", nil, &obj)
  // ...

  log.Println("Current user", currentUser)
}

```
For more examples, see package tests.

For list of supported/implemented endpoints, see [Endpoints.md](./endpoints.md)

### Pagination ###

All requests for resource collections (posts, pages, media, revisions, etc.)
support pagination. Pagination options are described in the
`wordpress.ListOptions` struct and passed to the list methods directly or as an
embedded type of a more specific list options struct (for example
`wordpress.PostsListOptions`). Pages information is available via the
`wordpress.Response` struct.

```go
client := wordpress.NewClient(&wordpress.Options{
  BaseAPIURL: API_BASE_URL, // example: `http://192.168.99.100:32777/wp-json/`
  Username:   USER,
  Password:   PASSWORD,
}, nil)

ctx := context.Background()

opt := &wordpress.PostsByOrgOptions{
  ListOptions: wordpress.ListOptions{PerPage: 10},
}
// get all pages of results
var allPosts []*wordpress.Post
for {
  posts, resp, err := client.Posts.List(ctx, opt)
  if err != nil {
    return err
  }
  allPosts = append(allPosts, posts...)
  if resp.NextPage == 0 {
    break
  }
  opt.Page = resp.NextPage
}
```

## Test
__Note:__
Before running the tests, ensure that you have set up your test environment


### Prerequisites
- Wordpress 4.x
- [WP-API's BasicAuth plugin (for authentication)](https://github.com/WP-API/Basic-Auth)

### Setting up test environment
- Install prequisites (see above)
- Import [./test-data/go-wordpress.wordpress.2015-08-23.xml](./test-data/go-wordpress.wordpress.2015-08-23.xml) to your local test Wordpress installation
- Upload at least one media to your Wordpress installation (Admin > Media > Upload)
- Edit one (1) most recent Post to create a revision
- Edit one (1) most recent Page to create a revision

## Running test


```bash

# Set test enviroment
export WP_API_URL=http://192.168.99.100:32777/wp-json/
export WP_USER=<user>
export WP_PASSWD=<password>

cd <path_to_package>/github.com/robbiet480/go-wordpress
go test

```

## TODO
- [ ] Support OAuth authentication (already supported by passing in your own HTTP client)
