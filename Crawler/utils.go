package crawler

import (
	"net/url"
	"path"
	"strings"
)

const (
	oldString = "https://swapi.co"
	newString = "http://localhost:8080/query"
)

// Combine api relative path and server address into a complete url
func composeUrl(api string) (string, error) {
	hostURL := url.URL{
		Scheme:   "https",
		Host:     "swapi.co",
		Path:     "api/" + api,
		RawQuery: "format=json",
	}

	return hostURL.String(), nil
}

// Replace the crawled server url with the URL of our server for string
func replaceSingleOldString(input *string) {
	*input = strings.Replace(*input, oldString, newString, 1)
}

// Replace the crawled server url with the URL of our server for []string
func replaceOldString(input []string) {
	for key := range input {
		input[key] = strings.Replace(input[key], oldString, newString, 1)
	}
}

// Extract id from link
func getIDFormUrl(u string) string {
	return path.Base(u)
}
