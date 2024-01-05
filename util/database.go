package util

import (
	"errors"
	"strings"
)

var (
	errNoScheme = errors.New("no scheme present in url.")
	errEmptyUrl = errors.New("empty url to parse scheme from.")
)

func SchemeFromURL(url string) (string, error) {
	if url == "" {
		return "", errEmptyUrl
	}

	i := strings.Index(url, ":")

	if i < 1 {
		return "", errNoScheme
	}

	return url[0:i], nil
}
