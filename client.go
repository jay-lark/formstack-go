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

func clientDo(so FormstackOptions, method string, path string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, constructUrl(path), bytes.NewReader(body))

	generateHeaders(so, req)

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

func generateHeaders(fo FormstackOptions, req *http.Request) {
	authToken := fmt.Sprintf("token %s", fo.ApiKey)
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", fo.UserAgent)
}

func constructUrl(path string) string {
	formstackEndpoint := "https://www.formstack.com/api/v2%s"
	return fmt.Sprintf(formstackEndpoint, path)
}
