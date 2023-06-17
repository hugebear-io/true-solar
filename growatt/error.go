package growatt

import (
	"net/http"

	"github.com/hugebear-io/gofiber/errors"
)

var (
	ErrResponseMustNotBeHTML = errors.NewInternalError("response must not be HTML")
	ErrTooManyRequests       = errors.NewError(http.StatusTooManyRequests, "too many request")
)
