# go-wp-api
Golang client library for WP-API (Wordpress REST API)


## Installation

```bash
go get github.com/sogko/go-wordpress

```

## Usage

### Quick example
```go
package main

import (
	"github.com/sogko/go-wordpress"
	"net/http"
)

func main() {

  // create wp-api client
  client := wordpress.NewClient(&wordpress.Options{
    BaseAPIURL: API_BASE_URL, // example: `http://192.168.99.100:32777/wp-json/wp/v2`
    Username:   USER,
    Password:   PASSWORD,
  })
  	
  // for eg, to get current user (GET /users/me)
  currentUser, resp, body, _ := client.Users().Me()
  if resp.StatusCode != http.StatusOK {
    // handle error
  }
  
  // `body` will contain raw JSON body in []bytes
  
  // Or you can use your own structs (for custom endpoints, for example)
  // Below is the equivalent of `client.Posts().Get(100, nil)`
  var obj MyCustomPostStruct
  resp, body, err := client.Get("/posts/100", nil, &obj)
  // ...
  
  log.Println("Current user", currentUser)
}

```
For more examples, see package tests.

For list of supported/implemented endpoints, see [Endpoints.md](./endpoints.md)


## Test
__Note:__
Before running the tests, ensure that you have set up your test environment


### Prerequisites
- Wordpress 4.x
- WP-API plugin
- WP-API's BasicAuth plugin (for authentication)
- [WP REST API Meta Endpoints plugin](https://github.com/WP-API/wp-api-meta-endpoints) (for Meta endpoints)

### Setting up test environment
- Install prequisits (see above)
- Import [./test-data/go-wordpress.wordpress.2015-08-23.xml](./test-data/go-wordpress.wordpress.2015-08-23.xml) to your local test Wordpress installation
- Upload at least one media to your Wordpress installation (Admin > Media > Upload)
- Edit one (1) most recent Post to create a revision
- Edit one (1) most recent Page to create a revision

## Running test


```bash

# Set test enviroment
export WP_API_URL=http://192.168.99.100:32777/wp-json/wp/v2
export WP_USER=<user>
export WP_PASSWD=<password>

cd <path_to_package>/github.com/sogko/go-wordpress
go test

```

## TODO
- [ ] `godoc` documentation, so its easier for library users to map the REST APIs to library calls 
- [ ] Test `comments` API endpoint. (Currently, already implemented but not tested due to WP-API issues with creating comments reliably)
- [ ] Support OAuth authentication
