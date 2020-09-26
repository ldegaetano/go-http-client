# Go http client

A Rest Client for Go (Golang)

An extremely simple to use and lightweight REST Client.

### Usage

````
go get "github.com/ldegaetano/go-http-client"
````

````
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ldegaetano/go-http-client/rest"
)

type Item struct {
	Code  string  `json:item_code"`
	Price float64 `json:item_price"`
}

func main() {
	cfg := rest.ClientCfg{
		RetryAttempts: 3,
		RetryInterval: time.Second,
		BasePath:      "/api/items",
		Timeout:       3 * time.Second,
	}
	client := rest.NewClient(cfg)

	res, body, err := client.Get("/prices")
	if err != nil {
		//Do something
	}

	if res.StatusCode == http.StatusInternalServerError {
		//Do somethin
	}
	fmt.Printf("GET response: %s", string(body))

	b, _ := json.Marshal(Item{
		Code:  "p1",
		Price: 10.7,
	})

	h := &http.Header{
		"X-application-id": []string{"54378784"},
	}
	res, body, err = client.PostWithHeader("/prices", b, h)
	if err != nil {
		//Do something
	}

	if res.StatusCode == http.StatusInternalServerError {
		//Do somethin
	}
	fmt.Printf("POST response: %s", string(body))
}
````

