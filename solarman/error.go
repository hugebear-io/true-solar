package solarman

import "github.com/hugebear-io/gofiber/errors"

var (
	ErrResponseMustNotBeHTML = errors.NewInternalError("response must not be HTML")
)
