# http/net package
HTTP or Hypertext Transfer Protocol is an application layer protocol. When you write a URL address in format of `https://website-address`, your browser is acting as a client and uses Hypertext Transfer Protocol to ask a server for resources and uses those resources to view the webpage.
In this concept we will learn how to code HTTP clients and servers using http/net package.

## Client
A client can be defined by using `http.Client` structure.
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

~~~~exercism/caution
Tip: You should always close the response body. Since Go uses HTTP/1.1, it keeps a connection to the server open in case you might reuse it. Therefore using `defer resp.Body.Close()` you can make sure the response body will be closed after your function exits.  
~~~~

### Head method:
`func (c *Client) Head(url string) (resp *Response, err error)`
This method lets you send HEAD requests to a URL and receive the response or any error that might occur. The result of this method is similar to Get, but the response body is empty.

~~~~exercism/note
You still need to close the response body even if you are sending a HEAD request.  
~~~~

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
A client can be defined by using `http.Server` structure. There are five important fields in this structure:
1. `Addr string`: The TCP address that the server will listen on,
2. `Handler Handler`: the server’s multiplexer,
3. `IdleTimeout time.Duration`: This is the maximum amount of time that a connection will be kept open before a new request from the client arrives.
4. `ReadHeaderTimeout time.Duration`: This is the maximum amount of time that the server bears to read the request headers.
5. `ReadTimeout time.Duration`: This is the maximum amount of time that the server bears to read the request. The core components used to create HTTP servers using `net/http` package are handlers.

~~~~exercism/note
As stated in [the server.go definition][server.go] regarding how new incoming http connections are handled:
"Serve accepts incoming HTTP connections on the listener l, creating a new service goroutine for each. The service goroutines read requests and then call handler to reply to them.". Therefore each connection has its own goroutine which enables go to handle client requests concurrently. Albeit, this leads to better performance especially for handling a large set of clients, you must be aware of shared access to a resource and use mutexes or channels to handle it.
~~~~

### Handlers:
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
Handlers are responsible for satisfying the client's requests by performing various tasks. Each object that implements the `http.Handler` interface, is capable of responding to client requests. Handler functions receive an `http.Request` which contains information regarding the client’s request, and an `http.ResponseWriter` that allows the handler to write the appropriate response in accordance with the request. `http.ResponseWriter` interface contains a `Header` field (!!! I need to complete this !!!), a `Write([]byte) (int, error)` method that writes the response to the client, and a `WriteHeader(statusCode int)`  that writes the appropriate status code to the client.

~~~~exercism/caution
If you wish to send and status code besides `http.StatusOk` you need to use `WriteHeader` before `Write`, since `Write` assumes `http.StatusOk` if no status is set before invoking it.  
Example:  
```go
func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Not Found!"))
	w.WriteHeader(http.StatusNotFound)
}
// The result of running this function is a response with status code 200 Ok!
----------
func(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found!"))
}
// The result of running this function is a response with status code 404 NotFound!
```
~~~~

### Multiplexer:
`http.ServeMux` is responsible for routing each request to the right handler by matching the pattern in the request with the **longest** pattern specified for it. `http.NewServeMux() *ServeMux` is used to create a new multiplexer. It uses handlers to route each request to the proper handler. As an example:
```go
func multiplexer() *ServeMux{
	serveMux := http.NewServeMux()
	serveMux := http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!"))
	})
	return serveMux
}
```

#### Absolute paths vs subtrees:  
Paths that end without a `/` are called absolute paths (e.g. `"root/image"`). If an absolute path is used in the request, the multiplexer will try to find the “exact” match for that path. On the contrary, paths that end with `/` are called subtrees (e.g. `“root/image/“`) and the multiplexer will try to find the best match for these paths. Beware that the multiplexer will convert an absolute path to a subtree if it does not find a corresponding handler.  

~~~~exercism/caution
Draining the request body before closing it:  
Unlike the client that drains the body before closing it, in the server you need to manually drain the request body. A good way to do so is writing a middleware that closes the request after the server job is done (I have seen this pattern for the first time in [Adam Woodbeck's mux_test.go][awoodbeck-gnp-chapter09-mux_test.go]):   
```go
func drainAndClose(next *ServeMux) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			_, _ = io.Copy(ioutil.Discard, r.Body)
			_ = r.Body.Close()
		},
	)
}

func multiplexer() http.Handler{
	serveMux := http.NewServeMux()
	serveMux := http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!"))
	})
	mux := drainAndClose(serveMux)
	return serveMux
}
```
~~~~

[awoodbeck-gnp-chapter09-mux_test.go]: https://github.com/awoodbeck/gnp/blob/master/ch09/mux_test.go
[server.go]: https://go.dev/src/net/http/server.go