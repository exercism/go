# Go net/http package server

[HTTP (Hypertext Transfer Protocol)](https://www.cloudflare.com/learning/ddos/glossary/hypertext-transfer-protocol-http/) is a protocol that allows devices to send and receive data using a client-server model.
In this concept we will learn how to code HTTP servers and respond to client requests using go `net/http` package.

## Server

A server can be defined by using `http.Server` structure. There are five important fields in this structure:

1. `Addr string`: The TCP address that the server will listen on,
2. `Handler Handler`: the server’s multiplexer,
3. `IdleTimeout time.Duration`: This is the maximum amount of time that a connection will be kept open before a new request from the client arrives.
4. `ReadHeaderTimeout time.Duration`: This is the maximum amount of time that the server bears to read the request headers.
5. `ReadTimeout time.Duration`: This is the maximum amount of time that the server bears to read the request.

~~~~exercism/note
If you include no timeout for your server, then faulty clients can keep using resources and eventually bring your server to a halt.
~~~~

The core components used to create HTTP servers using `net/http` package are handlers. Consider the example below:

```go
func main(){
    // code ...
    server := http.Server{
        Addr: "127.0.0.1:9898",
        Handler: myHandler,
        IdleTimeout: 30 * time.Second,
        ReadHeaderTimeout: 10 * time.Second,
        ReadTimeout: 20 * time.Second,
        }
    server.ListenAndServe() // Starting the server
    // code ...
}
```

~~~~exercism/note
As stated in [the server.go definition][server.go] regarding how new incoming http connections are handled: "Serve accepts incoming HTTP connections on the listener l, creating a new service goroutine for each. The service goroutines read requests and then call handler to reply to them.". Therefore each connection has its own goroutine which enables go to handle client requests concurrently. Albeit, this leads to better performance especially for handling a large set of clients, you must be aware of shared access to a resource and use mutexes or channels to handle it.
~~~~

### Handlers

Handlers are responsible for satisfying the client's requests by performing various tasks. Each object that implements the `http.Handler` interface, is capable of responding to client requests.

```go
type Handler interface {
 ServeHTTP(ResponseWriter, *Request)
}
```

Handler functions receive an `http.Request` which contains information regarding the client’s request, and an `http.ResponseWriter` that allows the handler to write the appropriate response in accordance with the request. `http.ResponseWriter` interface contains a `Header` field, a `Write([]byte) (int, error)` method that writes the response to the client, and a `WriteHeader(statusCode int)` that writes the appropriate status code to the client.

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

### Multiplexer

`http.ServeMux` is responsible for routing each request to the right handler by matching the pattern in the URL of the request with the **longest** pattern specified for it. `http.NewServeMux() *ServeMux` is used to create a new multiplexer. As an example:

```go
func multiplexer() *ServeMux{
 serveMux := http.NewServeMux()
 serveMux := http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
 w.Write([]byte("Hello World!"))
 })
 return serveMux
}
```

Note that multiplexers satisfy handler's interface, therefore to use one, we can initialize a server by assigning a multiplexer as its handler.

#### Absolute paths vs subtrees

Paths that end without a `/` are called absolute paths (e.g. `"root/image"`). If an absolute path is used in the request, the multiplexer will try to find the “exact” match for that path. On the contrary, paths that end with `/` are called subtrees (e.g. `“root/image/“`) and the multiplexer will try to find the best match for these paths. Beware that the multiplexer will convert an absolute path to a subtree if it does not find a corresponding handler.  

~~~~exercism/caution
Draining the request body before closing it:  
Unlike the client that drains the body before closing it, in the server you need to manually drain the request body. A good way to do so is writing a middleware that closes the request after the server job is done (I have seen this pattern for the first time in [Adam Woodbeck's mux_test.go][awoodbeck-gnp-chapter09-mux_test.go]):   

```go
func drainAndClose(next *http.ServeMux) http.Handler {
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
 return mux
}
```

~~~~

[awoodbeck-gnp-chapter09-mux_test.go]: https://github.com/awoodbeck/gnp/blob/master/ch09/mux_test.go
[server.go]: https://go.dev/src/net/http/server.go