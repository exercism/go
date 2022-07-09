# net/http exercise

Goblinocus president was happy with your previous work and now he demands your return to complete the weather forecast project! You are required to write an HTTP server that can respond to requests regarding the weather forecast of various cities in Goblinocus and an HTTP client that residents can use to get the weather forecast of their hometowns.

## 1. Server

ForecastServer must return a server that can appropriately handle clients' requests. You need to have a multiplexer that routes each request to its appropriate handler:

1. if address ends with `“/city”`, it will send the request to a handler that sends back the forecast of the requested weather. Your handler must support both `http.Get` and `http.Post` methods. For any other method respond with MethodNotAllowed status code. If the requested city is incorrect, respond with a NotFound status code.
2. For any other path, you need to use a handler to send back a BadRequest status code.
3. You can get available list of cities by calling `Cities()` and get each cities' forecast by calling `Forecast(city)`.

## 2. Client

The client is responsible to connect to the server via `ForecastClient` function, ask for a specific city’s weather and return the result with any possible error that have occured.
