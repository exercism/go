package weatherforecastclient

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func ForecastClient(serverAddr, method string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, serverAddr, http.NoBody)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
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
}
