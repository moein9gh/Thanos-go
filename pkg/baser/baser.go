package baser

import (
	"encoding/base64"
	"strings"
)

func DecodeString(src string) ([]byte, error) {
	var isRaw = !strings.HasSuffix(src, "=")
	if strings.ContainsRune(src, '/') {
		if isRaw {
			return base64.RawStdEncoding.DecodeString(src)
		}
		return base64.StdEncoding.DecodeString(src)
	}

	if isRaw {
		return base64.RawURLEncoding.DecodeString(src)
	}
	return base64.URLEncoding.DecodeString(src)
}

func EncodeToString(src []byte) string {
	return base64.RawURLEncoding.EncodeToString(src)
}

func EncodeToStringStd(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
