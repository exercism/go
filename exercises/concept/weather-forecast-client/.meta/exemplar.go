package weatherforecastclient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	serverAddr string
}

func NewForecastClient(serverAddr string) (*Client, error) {
	return &Client{httpClient: &http.Client{Timeout: 5 * time.Second}, serverAddr: serverAddr}, nil
}

// SELF:Check what is the standard way to structure the Get and Post different functions
func (c *Client) Query(city, method string, data ...string) (string, error) {
	if method == http.MethodGet {
		req, err := http.NewRequest(method, fmt.Sprintf("%s/weather?city=%s", c.serverAddr, city), http.NoBody)
		if err != nil {
			return "", err
		}
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return "", errors.New(fmt.Sprint(resp.StatusCode))
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(b), nil
	} else if method == http.MethodPost {
		// SELF: complete this section
	}
	// SELF: find a better way to raise this error
	return "", errors.New(fmt.Sprint(http.StatusMethodNotAllowed))
}
