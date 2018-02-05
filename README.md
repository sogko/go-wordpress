# go-wordpress
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
  "context"
  "log"
  "net/http"

  "github.com/robbiet480/go-wordpress"
)

func main() {

  tp := wordpress.BasicAuthTransport{
    Username: USER,
    Password: PASSWORD,
  }

  // create wp-api client
  client, _ := wordpress.NewClient(API_BASE_URL, tp.Client())

  ctx := context.Background()

  // for eg, to get current user (GET /users/me)
  currentUser, resp, _ := client.Users.Me(ctx, nil)
  if resp != nil && resp.StatusCode != http.StatusOK {
    // handle error
  }

  // Or you can use your own structs (for custom endpoints, for example)
  // Below is the equivalent of `client.Posts.Get(ctx, 100, nil)`
  var obj MyCustomPostStruct
  resp, err := client.Get(ctx, "/posts/100", nil, &obj)
  // ...

  fmt.Printf("Current user %+v", currentUser)
}
```

For more examples, see package tests.

For list of supported/implemented endpoints, see [Endpoints.md](./endpoints.md)

### Authentication

The go-wordpress library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you.

Note that when using an authenticated Client, all calls made by the client will
include the specified authentication transport token. Therefore, authenticated clients should
almost never be shared between different users.

#### Username/Password or Application Password

A basic authentication (username/password) client for use with
the [WP-API BasicAuth plugin](https://github.com/WP-API/Basic-Auth)
or [Application Passwords plugin](https://wordpress.org/plugins/application-passwords/)
is included with the library.
An example implementation can be found in [example/basicauth/main.go](example/basicauth/main.go).

#### OAuth 1.0a

If you use the [OAuth 1.0a Server](https://github.com/WP-API/OAuth1) for authentication,
you can find an example implementation in [example/oauth2/main.go](example/oauth2/main.go) using the oauth1 library
(which is very similar to the official OAuth 2.0 library).
See the [oauth1 docs](https://godoc.org/dghubble/oauth1) for complete instructions on using that library.

#### OAuth 2.0 and JWT

If you are using the [JWT](https://wordpress.org/plugins/jwt-authentication-for-wp-rest-api/) plug-in for authentication,
you can use the [oauth2](https://github.com/golang/oauth2) library's `StaticTokenSource`.
An example implementation can be found in [example/oauth2/main.go](example/oauth2/main.go).
See the [oauth2 docs](https://godoc.org/golang.org/x/oauth2) for complete instructions on using that library.

#### Other authentication styles

For any other authentication methods, you should only need to provide a custom `http.Client` when creating a new WordPress client.

### Pagination

All requests for resource collections (posts, pages, media, revisions, etc.)
support pagination. Pagination options are described in the
`wordpress.ListOptions` struct and passed to the list methods directly or as an
embedded type of a more specific list options struct (for example
`wordpress.PostListOptions`). Pages information is available via the
`wordpress.Response` struct.

```go
package main

import (
  "context"

  "github.com/robbiet480/go-wordpress"
)

func main() {
  tp := wordpress.BasicAuthTransport{
    Username: USER,
    Password: PASSWORD,
  }

  // create wp-api client
  client, _ := wordpress.NewClient(API_BASE_URL, tp.Client())

  ctx := context.Background()

  opt := &wordpress.PostListOptions{
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

## Running tests

```bash
# Set test enviroment
export WP_API_URL=http://192.168.99.100:32777/wp-json/
export WP_USER=<user>
export WP_PASSWD=<password>

cd $GOPATH/src/github.com/robbiet480/go-wordpress
go test
```

## Thanks

Large parts of this library were inspired if not outright copied from Google's excellent [`go-github`](https://github.com/google/go-github) library.
