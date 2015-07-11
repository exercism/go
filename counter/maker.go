package counter

import (
	"log"
	"os"
)

// A little trick to help test the various versions,
// set the COUNTER_IMPL environment variable to the
// number of the implementation you want to test.
// E.g. run `COUNTER_IMPL=4 go test` to get Impl4.

func makeCounter() Counter {
	switch os.Getenv("COUNTER_IMPL") {
	case "1":
		return &Impl1{}
	case "2":
		return &Impl2{}
	case "3":
		return &Impl3{}
	case "4":
		return &Impl4{}
	case "":
		log.Fatalf("Don't forget to set COUNTER_IMPL")
	default:
		log.Fatalf("Unknown COUNTER_IMPL value: %s", os.Getenv("COUNTER_IMPL"))
	}
	panic("not reachable, but go's return analysis needs it")
}
