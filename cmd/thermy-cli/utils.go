package main

import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func createRequest(method, path string, queryParams url.Values, body any) (*http.Request, error) {
	reqUrl, err := url.Parse(appUrl)
	if err != nil {
		return nil, fmt.Errorf("client: could not parse url: %s", appUrl)
	}
	reqUrl.Path = apiPrefix + path
	reqUrl.RawQuery = queryParams.Encode()

	var reqBody io.ReadWriter
	switch v := body.(type) {
	case []byte:
		reqBody = bytes.NewBuffer(v)
	default:
		reqBody = new(bytes.Buffer)
		err = json.NewEncoder(reqBody).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("client: could not encode request body: %w", err)
		}
	}

	req, err := http.NewRequest(method, reqUrl.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("client: could not create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func doRequest(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client: error making http request: %w", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("client: could not read the response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("client: unexpected http status: %s (response body: %s)", res.Status, string(body))
	}
	ct := res.Header.Get("Content-Type")
	if ct != "" && ct != "application/json; charset=utf-8" {
		return nil, fmt.Errorf("client: unexpected content-type: %s", ct)
	}

	return body, nil
}
