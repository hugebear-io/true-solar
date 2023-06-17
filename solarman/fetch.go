package solarman

import (
	"io"
	"net/http"

	"github.com/hugebear-io/gofiber/fabric"
)

func prepareHttpResponse[R interface{}](req *http.Request) (*R, int, error) {
	// request to endpoint
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	// read a bytes of data
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// check empty response body
	if len(resBody) == 0 {
		return nil, res.StatusCode, nil
	}

	// decode response body
	var result R
	if err := fabric.Recast(resBody, &result); err != nil {
		if err := checkHTMLResponse(resBody); err != nil {
			return nil, http.StatusInternalServerError, err
		}

		return nil, http.StatusInternalServerError, err
	}

	return &result, res.StatusCode, nil
}
