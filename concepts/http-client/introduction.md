# HTTP

[HTTP (Hypertext Transfer Protocol)](https://www.cloudflare.com/learning/ddos/glossary/hypertext-transfer-protocol-http/) is a protocol that allows devices to send and receive data using a client-server model.
In this concept we will learn how to code HTTP clients and send requests to http servers using go `net/http` package.

## Setting up an HTTP client

A client can be defined by using `http.Client` structure:

```go
func NewClient() {
    client := &http.Client{}
}
```

### Sending a GET request

`func (c *Client) Get(url string) (resp *Response, err error)`
This method lets you send GET requests to a URL and receive the response or any error that might occur. Consider the example below:

```go
func client(url string) {
	client := &http.Client{}
	resp, err := client.Get(url) // Sending a GET request to google
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()            // Always close the reponse body after checking for any errors
	body, err := io.ReadAll(resp.Body) // Reading the response HTTP body
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body)) // Printing out the response body
}
```

Here we connect to Google main webpage and display the response body.

~~~~exercism/caution
Since Go uses HTTP/1.1, it keeps each connection to the server open in case you might reuse it. Go won't try to figure out if you are done with the connection and as a result, if you do not close the connection after you are done with it, it leads to memory leaks. Therefore it is extremely important to use `defer resp.Body.Close()` so you make sure the response body will be closed after your function exits. 
~~~~

### Sending a HEAD request

`func (c *Client) Head(url string) (resp *Response, err error)`
This method lets you send HEAD requests to a URL and receive the response or any error that might occur. The result of this method is similar to Get, but the response body is empty.

~~~~exercism/note
You still need to close the response body even if you are sending a HEAD request.  
~~~~

### Sending a POST request

`func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)`
This method lets you send a POST request to a URL and ask the server to accept the data you send in `body` with `contentType` type. Consider the example below:

```go
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
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusAccepted { // Checking whether the server accepted the request
		log.Fatalf("Server did not accept the request. Status code: %d", resp.StatusCode)
	}
}
```

Here we send a `"Hello, World!"` message, as type `"text/plain"` to a url given as input to `client` function.

### Do method

`func (c *Client) Do(req *Request) (*Response, error)`
Using `Do` method, you can send `GET`, `POST`, and `HEAD` requests. This method takes an `*http.Request` and sends a request to the server according to the variables set in the request. Consider the example below as a substitue for above examples:

```go
func client(url, method string) {
	client := &http.Client{}
	var body bytes.Buffer
	if method == http.MethodPost {
		enc := gob.NewEncoder(&body)
		_ = enc.Encode("Hello, World!")
	}

	// Creating the http.request
	req, err := http.NewRequest(method, url, &body)

	// Sending the request using client.Do
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted { // Checking whether the server accepted the request
		log.Fatalf("Server did not accept the request. Status code: %d", resp.StatusCode)
	}
}
```

### Timeout variable

`Timeout time.Duration`
`Client` type in Go has a variable called `Timeout`. This variable is responsible for closing the connection if the client's connection time, redirect, and response body read takes more time that the designated value. The default value for `Timeout` is zero, which means no timeout. This might cause unwanted behavior and a malicious server or unintended bug, can cause your program to halt. Therefore you should set a value for this variable every time that you create a new client to prevent this from happening. Here is a how you can do it while creating a client:

```go
func NewClient() *http.Client{
    return &http.Client{Timeout: 5 * time.second}
}
```