# http/net package
HTTP or Hypertext Transfer Protocol is an application layer[^1] protocol. When you write a URL address in format of `https://website-address`, your browser is acting as a client and uses Hypertext Transfer Protocol to ask a server for resources and uses those resources to view the webpage.
In this concept we will learn how to code HTTP clients and servers using http/net package.

## Client
There are three main methods that you can use in http/net package to act as an HTTP client:

### Get method:
`func (c *Client) Get(url string) (resp *Response, err error)`
This method lets you send GET requests to a URL and receive the response or any error that might occur. Consider the example below:
```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://www.google.com/") // Sending a GET request to google
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // Always close the reponse body

	body, err := io.ReadAll(resp.Body) // Reading the response HTTP body
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body)) // Printing out the response body
}
```
Here we connect to Google main webpage and display the response body.
> Tip: You should always close the response body. Since Go uses HTTP/1.1, it keeps a connection to the server open in case you might reuse it. Therefore using `defer resp.Body.Close()` you can make sure the response body will be closed after your function exits.

### Head method:
`func (c *Client) Head(url string) (resp *Response, err error)`
This method lets you send HEAD requests to a URL and receive the response or any error that might occur. The result of this method is similar to Get, but the response body is empty.
> Tip: You still need to close the response body even if you are sending a HEAD request.

### Post method:
`func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)`
This method lets you send a POST request to a URL and ask the server to accept the data you send in `body` with `contentType` type. Consider the example below:
```go
package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"
)

func client(url string) {
	// Encoding out message to send to server
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Post(url, "text/plain", &buff) // Trying to send the message
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusAccepted { // Checking whether the server accepted the request
		log.Fatalf("Server did not accept the request. Status code: %d", resp.StatusCode)
	}
	_ = resp.Body.Close()
}
```
Here we send a `"Hello, World!"` message, as type `"text/plain"` to a url given as input to `client` function. 

## Server