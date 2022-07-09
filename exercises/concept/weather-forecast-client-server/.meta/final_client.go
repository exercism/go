package serverAns

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ForecastClient(city, serverAddr, method string) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, serverAddr, strings.NewReader(city))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprint(resp.StatusCode))
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
