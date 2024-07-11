package base64

import (
	"encoding/base64"
)

func EncodeBase64(src []byte) string {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(src)))

	return string(dst)
}
