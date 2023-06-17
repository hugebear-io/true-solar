package growatt

import (
	"fmt"
	"io"
	"net/http"

	"github.com/hugebear-io/gofiber/errors"
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

	if res.StatusCode == http.StatusTooManyRequests {
		return nil, http.StatusTooManyRequests, ErrTooManyRequests
	}

	// read a bytes of data
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := checkHTMLResponse(resBody); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// check empty response body
	if len(resBody) == 0 {
		return nil, res.StatusCode, nil
	}

	// decode response body
	var result R
	if err := fabric.Recast(resBody, &result); err != nil {
		errRespPointer := APIErrorResponse{}
		if err := fabric.Recast(resBody, &errRespPointer); err != nil {
			return nil, http.StatusInternalServerError, err
		} else {
			errResp := errRespPointer.Result()
			growattErr := errors.NewError(http.StatusBadRequest, fmt.Sprintf("[%v]: %v", errResp.ErrorCode, errResp.ErrorMsg))
			return nil, http.StatusBadRequest, growattErr
		}
	}

	return &result, res.StatusCode, nil

}
