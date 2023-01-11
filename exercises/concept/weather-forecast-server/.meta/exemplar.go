package weatherforecastserver

import (
	"net/http"
	"time"
)

// import (
// 	"io"
// 	"net/http"
// 	"time"
// )

// type WeatherServer struct {
// 	Mux         *http.ServeMux
// 	BaseURL     string
// 	IdleTimeout time.Duration
// }

// func NewWeatherServer(url string, idleTimeout time.Duration) *WeatherServer {
// 	return &WeatherServer{
// 		Mux:         multiplexer(),
// 		BaseURL:     url,
// 		IdleTimeout: idleTimeout,
// 	}
// }

// func (ws *WeatherServer) Handle(path string, handler http.Handler) {
// 	ws.Mux.Handle(path, handler)
// }

// func (ws *WeatherServer) multiplexer() http.Handler {
// 	serveMux := http.NewServeMux()
// 	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
// 		writer.WriteHeader(http.StatusBadRequest)
// 		writer.Write([]byte("Please use '/city' to ask for weather status."))
// 	})
// 	serveMux.HandleFunc("/city", func(writer http.ResponseWriter, request *http.Request) {
// 		if request.Method == http.MethodGet {
// 			writer.Write([]byte("Goblinocus will have a nice sunny day tomorrow!"))
// 		} else {
// 			writer.WriteHeader(http.StatusMethodNotAllowed)
// 		}
// 	})
// 	mux := drainAndClose(serveMux)
// 	return mux
// }

// func drainAndClose(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		next.ServeHTTP(writer, request)
// 		_, _ = io.Copy(io.Discard, request.Body)
// 		_ = request.Body.Close()
// 	})
// }

type WeatherServer struct {
	httpServer *http.Server
	baseURL string
	idleTimeout time.Duration
	mux http.ServeMux
}

func NewWeatherServer(baseURL string, idleTimeout time.Duration) WeatherServer{
	server := &http.Server{
		Addr: baseURL,
		IdleTimeout: idleTimeout,
		Handler: ,
	}
}

func (ws *WeatherServer) Handle(path string, handler http.Handler) {
	
}

func multiplexer() http.Handler {
	serveMux := http.NewServeMux()
	return drainAndClose(serveMux)
}

func drainAndClose(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)
		_, _ = io.Copy(io.Discard, request.Body)
		_ = request.Body.Close()
	})
}