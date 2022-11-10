package weatherforecastclient

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type WeatherClient struct {
	httpClient *http.Client
	baseURL    string
	timeout    time.Duration
}

func NewWeatherClient(URL string, timeout time.Duration) *WeatherClient {
	client := &http.Client{
		Timeout: timeout,
	}
	return &WeatherClient{
		httpClient: client,
		baseURL:    URL,
		timeout:    timeout,
	}
}

func (wc *WeatherClient) TodayWeatherClient(city string) (*http.Response, error) {
	queryURL := fmt.Sprintf("%s/?city=%s", wc.baseURL, city)
	req, err := http.NewRequest(http.MethodGet, queryURL, http.NoBody)
	if err != nil {
		return nil, err
	}
	resp, err := wc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *WeatherClient) FutureWeatherQuery(city string, futureDate int) (*http.Response, error) {
	queryURL := fmt.Sprintf("%s/?city=%s", wc.baseURL, city)
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(futureDate)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, queryURL, &buff)
	if err != nil {
		return nil, err
	}
	resp, err := wc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *WeatherClient) WeatherReader(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}
