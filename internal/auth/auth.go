package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPUKey extracts the API key from the header of an HTTP request
//Example of usage:
// Authorization: ApiKey {insert API key here}

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}