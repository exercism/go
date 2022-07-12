# HTTP server

Goblinocus president was happy with your previous work and now he demands your return to complete the weather forecast project! You are required to write an HTTP server that can respond to requests regarding the weather forecast for Goblinocus and an HTTP client that residents can use it.

## Server

ForecastServer must return a server that can appropriately handle clients' requests. You need to have a multiplexer that routes each request to its appropriate handler:

1. if address ends with `“/city”`, it will send the request to a handler that sends back the forecast of the requested weather. Your handler must support only `http.Get` method. For any other method respond with MethodNotAllowed status code.
2. For any other path, you need to use a handler to send back a BadRequest status code.
