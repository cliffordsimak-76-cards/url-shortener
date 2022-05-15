package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func StringToMD5(str string) string {
	h := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", h[:8])
}

func MakeResultString(baseURL string, shortURL string) string {
	return strings.Join([]string{baseURL, shortURL}, "/")
}
