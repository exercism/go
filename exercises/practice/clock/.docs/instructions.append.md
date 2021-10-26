To satisfy the README requirement about clocks being equal, values of
your Clock type need to work with the == operator. This means that if your
New function returns a pointer rather than a value, your clocks will
probably not work with ==.

While the time.Time type in the standard library (https://golang.org/pkg/time/#Time)
doesn't necessarily need to be used as a basis for your Clock type, it might
help to look at how constructors there (Date and Now) return values rather
than pointers. Note also how most time.Time methods have value receivers
rather than pointer receivers.

For some useful guidelines on when to use a value receiver or a pointer
receiver see: https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
