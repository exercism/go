package weatherforecast

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestForecastClient(t *testing.T) {
	testCases := []struct {
		serverAddr  string
		method      string
		expected    string
		errExpected bool
	}{
		{fmt.Sprintf("http://%s/", addr), http.MethodPost, "", true},
		{fmt.Sprintf("http://%s/City", addr), http.MethodPut, "", true},
		{fmt.Sprintf("http://%s/city", addr), http.MethodGet, "Goblinocus will have a nice sunny day tomorrow!", false},
	}
	srv := tForecastServer(addr, 10*time.Second)
	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	for i, c := range testCases {
		resp, err := ForecastClient(c.serverAddr, c.method)
		if (err == nil) && c.errExpected {
			t.Fatalf("expected error, returned nil. test case number: %d", i+1)
		}
		if (err != nil) && !c.errExpected {
			t.Fatalf("did not expect an error. received: %s. test case number: %d", err, i+1)
		}
		if resp != c.expected {
			t.Fatalf("expected %s, received %s. test case number: %d", c.expected, resp, i+1)
		}
	}
	srv.Close()
}

func tForecastServer(addr string, idle time.Duration) http.Server {
	serv := http.Server{Addr: addr, Handler: tmultiplexer(), IdleTimeout: idle}
	return serv
}

func tmultiplexer() http.Handler {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Please use '/city' to ask for weather status."))
	})
	serveMux.HandleFunc("/city", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			writer.Write([]byte("Golinocus will have a nice sunny day tomorrow!"))
		} else {
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux := tdrainAndClose(serveMux)
	return mux
}

func tdrainAndClose(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)
		_, _ = io.Copy(io.Discard, request.Body)
		_ = request.Body.Close()
	})
}
