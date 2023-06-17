package growatt

import "regexp"

var htmlTagsRegExp = regexp.MustCompile(`<\/?[a-z][\s\S]*>`)

const (
	URL_VERSION1  = "https://openapi.growatt.com/v1"
	AUTH_HEADER   = "Token"
	MAX_PAGE_SIZE = 100
)
