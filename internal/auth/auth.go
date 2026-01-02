package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("api key not found")
	}

	apiKey := strings.Split(val, " ")

	if len(apiKey) != 2 || apiKey[0] != "ApiKey" {
		return "", errors.New("Invalid api key format")
	}

	return apiKey[1], nil
}
