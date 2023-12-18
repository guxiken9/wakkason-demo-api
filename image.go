package main

import (
	"encoding/base64"
	"errors"
	"strings"
)

func DecodeBase64(text string) ([]byte, error) {

	t := strings.Split(text, ",")
	if len(t) == 2 {
		result, err := base64.StdEncoding.DecodeString(t[1])
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, errors.New("不正なBase64データが指定されましたあ")
}
