package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// ErrFrom constructs err from response data
func ErrFrom(rd *ResponseData) error {
	if rd.Success {
		return nil
	}
	return errors.New(rd.Message)
}

// ExtractData ...
func ExtractData(resp *http.Response) (*ResponseData, error) {
	rd := &ResponseData{}
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(rd)
	return rd, err
}

//RequestWithJSON payload
func RequestWithJSON(method, url string, data interface{}, token string) (*http.Response, error) {
	buf := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return nil, fmt.Errorf("when encoding request body: %v", err)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, fmt.Errorf("when creating request: %v", err)
	}

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := http.Client{}

	return client.Do(req)
}
