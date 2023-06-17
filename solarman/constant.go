package solarman

import "regexp"

var htmlTagsRegExp = regexp.MustCompile(`<\/?[a-z][\s\S]*>`)

const (
	URL_VERSION1         = "https://globalapi.solarmanpv.com"
	AUTHORIZATION_HEADER = "Authorization"
	MAX_PAGE_SIZE        = 200
)

const (
	TIME_TYPE_TIMEFRAME = iota + 1
	TIME_TYPE_DAY
	TIME_TYPE_MONTH
	TIME_TYPE_YEAR
)

const BRAND = "solarman"

const (
	DATA_LIST_KEY_CUMULATIVE_PRODUCTION = "Et_ge0"
	DATA_LIST_KEY_GENERATION            = "generation"
)

const (
	DEVICE_STATUS_ON      = "ONLINE"
	DEVICE_STATUS_OFF     = "OFFLINE"
	DEVICE_STATUS_FAILURE = "FAILURE"
)