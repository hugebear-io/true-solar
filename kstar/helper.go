package kstar

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func DecodePassword(password string) string {
	hashPassword := md5.Sum([]byte(password))
	return strings.ToUpper(fmt.Sprintf("%x", hashPassword[:]))
}
