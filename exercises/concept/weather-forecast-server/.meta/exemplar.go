package weatherforecastserver

import (
	"io"
	"net/http"
	"time"
)

func ForecastServer(addr string, idle time.Duration) http.Server {
	serv := http.Server{Addr: addr, Handler: multiplexer(), IdleTimeout: idle}
	return serv
}

func multiplexer() http.Handler {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Please use '/city' to ask for weather status."))
	})
	serveMux.HandleFunc("/city", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			writer.Write([]byte("Goblinocus will have a nice sunny day tomorrow!"))
		} else {
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux := drainAndClose(serveMux)
	return mux
}

func drainAndClose(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)
		_, _ = io.Copy(io.Discard, request.Body)
		_ = request.Body.Close()
	})
}
