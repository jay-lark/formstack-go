package formstack

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type FormstackOptions struct {
	ApiKey    string
	UserAgent string
}

var ErrInvalidAuthn = errors.New("credentials not valid")
var ErrInvalidAuthz = errors.New("credentials not authorized to access resource")
var ErrNotFound = errors.New("requested resource not found")
var ErrUnexpectedStatus = errors.New("unexpected HTTP status code")

func generateHeaders(fo FormstackOptions, req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+fo.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fo.UserAgent)
}

func constructUrl(path string) string {
	formstackEndpoint := "https://www.staging-formstack.com/api/v2" + path
	return formstackEndpoint
}

func clientDo(fo FormstackOptions, method string, path string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, constructUrl(path), bytes.NewReader(body))
	fmt.Println(req)
	generateHeaders(fo, req)
	fmt.Println(req)
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 300 {
		return res, nil
	} else if res.StatusCode == 401 {
		return nil, fmt.Errorf("%w", ErrInvalidAuthn)
	} else if res.StatusCode == 403 {
		return nil, fmt.Errorf("%w", ErrInvalidAuthz)
	} else if res.StatusCode == 404 {
		return nil, fmt.Errorf("%w", ErrNotFound)
	} else {
		return nil, errors.New(strconv.Itoa(res.StatusCode))
	}
}
