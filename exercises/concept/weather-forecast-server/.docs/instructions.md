# HTTP server

Goblinocus president was happy with your previous work and now he demands your return to complete the weather forecast project! You are required to write an HTTP server that can respond to requests regarding the weather forecast for Goblinocus and an HTTP client that residents can use it.

## Server

NewWeatherServer must return a server that can appropriately handle clients' requests. You need to have a multiplexer that routes each request to its appropriate handler. There are two types of requests that clients are allowed to make:

1. A GET request that asks for a city's weather for today. The URL must be in form of `"baseUrl/?city={cityName}"`. If the URL form is different, the request must be rejected (with `http.StatusBadRequest`). If the request is accepted, you must extract the city name from the queried URL and respond with "`{cityName} will have a nice weather today!`",
2. A POST request that asks for a city's weather in the future. The URL must be in form of `"baseUrl/?city={cityName}"`. If the URL form is different or the requested date is not positive, the request must be rejected (with `http.StatusBadRequest`). If the request is accepted, you must extract the city name from the queried URL, and the future date from the sent request body, and respond with "`{cityName} will have a nice weather in {futureDate} days!`".

Be wary, your clients might be faulty and take a lot of time to send a request, therefore implement measures to close a connection if it is unused for more than 5 seconds.
