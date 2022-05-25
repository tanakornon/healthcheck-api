package utils

import (
	"strings"
)

const (
	Http  string = "http://"
	Https        = "https://"
)

func ParseUrl(url string) string {
	hasHttp := strings.HasPrefix(url, Http)
	hasHttps := strings.HasPrefix(url, Https)

	if !hasHttp && !hasHttps {
		return Http + url
	}

	return url
}
