package base64

import (
	"encoding/base64"
)

func DecodeBase64(s string) (string, error) {
	sd, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		return "", e
	}
	return string(sd), nil
}
