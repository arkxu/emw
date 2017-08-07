package goutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAndFill get the JSON object from remote url and unmarshal to object
func GetAndFill(urlStr string, obj interface{}) error {

	resBytes, err := Get(urlStr)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resBytes, &obj); err != nil {
		return err
	}
	return nil
}

// Get get bytes from remote url
func Get(urlStr string) ([]byte, error) {

	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("http return code: %d", res.StatusCode)
	}

	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBytes, nil
}
