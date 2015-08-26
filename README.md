# go-wp-api
Golang client library for WP-API (Wordpress REST API)


## Installation

```bash
go get github.com/sogko/go-wordpress

```

## Usage
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
  currentUser, resp, _, _ := client.Users().Me()
  if resp.StatusCode != http.StatusOK {
    // handle error
  }
  
  log.Println("Current user", currentUser)
}

```
For more examples, see package tests.

For list of supported/implemented endpoints, see [Endpoints.md](./endpoints.md)

## Test

### Prerequisites
- Wordpress 4.x
- WP-API plugin
- WP-API's BasicAuth plugin (for authentication)

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
- [ ] Test `comments` API endpoint. (Currently, already implemented but not tested due to WP-API issues with creating comments reliably)
- [ ] Support OAuth authentication
