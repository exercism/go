# About

A [`Time`][time] in Go is a type describing a moment in time. The date and time information can be accessed, compared, and manipulated through its methods, but there are also some functions called on the `time` package itself. The current date and time can be retrieved through the [`time.Now`][now] function.

The [`time.Parse`][parse] function parses strings into `Time` types using the particular format string `Mon Jan 2 15:04:05 -0700 MST 2006`. More on how this works can be found [here][article].

The `time` package includes another type, `Duration`, representing elapsed time, plus support for locations/time zones, timers, and other related functionality that will be covered in another concept.

[time]: https://golang.org/pkg/time/#Time
[now]: https://golang.org/pkg/time/#Now
[parse]: https://golang.org/pkg/time/#Parse
[article]: https://www.pauladamsmith.com/blog/2011/05/go_time.html
