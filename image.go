package main

import "encoding/base64"

func DecodeBase64(text string) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}

	return result, err
}
