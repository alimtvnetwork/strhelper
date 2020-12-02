package strhashinternal

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) [16]byte {
	data := []byte(str)

	return md5.Sum(data)
}

func Md5Ptr(str *string) [16]byte {
	if str == nil {
		return Md5("")
	}

	data := []byte(*str)

	return md5.Sum(data)
}

func Md5HashString(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
