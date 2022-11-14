# HTTP client

Now that you have successfully created a weather server for the president, they need an HTTP client to be able to make queries and get responses. They have given you the following list of tasks to fulfill for your weather client:

## 1. Create a WeatherClient type

First, you need to create a Client type that must have three fields:

- A `*http.Client` to hold the main http client
- A `string` to hold the main server's base URL
- A `time.Duration` to hold the timeout value for the client

## 2. Create a function named NewWeatherClient

Now you must define a function that gets all the required arguments and returns a new WeatherClient with the designated parameters.

## 3. Create a method named TodayWeatherQuery

`todayWeatherQuery` gets a city name as its argument and returns the response. You must send a `GET` request and attach the name of the city to the main URL for making the query using the template below:
`"baseUrl/?city={cityName}"`

## 4. Create a method named FutureWeatherQuery

`futureWeatherQuery` gets a city name and an integer as its arguments and returns the response for the city's weather in the future. You must send a `POST` request, send the integer to get the weather for n-th day in the future, and use the URL rule above.

## 5. Create a method named WeatherReader

`WeatherReader` gets an *http.Response and an error as its arguments and returns the result as a string and any possible error encountered.
