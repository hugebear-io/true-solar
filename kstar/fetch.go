package kstar

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/hugebear-io/gofiber/errors"
	"github.com/hugebear-io/gofiber/fabric"
)

func prepareHttpRequest(method string, url string, values url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	contentType := "application/x-www-form-urlencoded"
	contentLength := strconv.Itoa(len(values.Encode()))
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", contentLength)

	return req, nil
}

func prepareHttpResponse[R interface{}](req *http.Request) (*R, int, error) {
	// request to endpoint
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		return nil, http.StatusTooManyRequests, ErrorTooManyRequest
	}

	// read a bytes of data
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// check empty response body
	if len(resBody) == 0 {
		return nil, res.StatusCode, nil
	}

	var defaultResponse APIDefaultResponse
	if err := fabric.Recast(resBody, &defaultResponse); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	meta := defaultResponse.Meta.Result()
	if !meta.Success {
		return nil, http.StatusInternalServerError, errors.NewInternalError(meta.Code)
	}

	// decode response body
	var result R
	if err := fabric.Recast(resBody, &result); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &result, res.StatusCode, nil

}
