package huawei

import "github.com/hugebear-io/gofiber/errors"

// TODO: map error message
var (
	ErrorTooManyRequest = errors.NewInternalError("too many request")
)
