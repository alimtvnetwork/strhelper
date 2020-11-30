package strhash

import (
	"crypto/md5"
	"encoding/hex"

	"gitlab.com/evatix-go/strhelper/anyto"
)

var (
	// nil or empty string hash is same https://play.golang.org/p/cdq61I5eO2g
	EmptyMd5Hash       = md5.Sum([]byte(""))
	EmptyMd5HashString = hex.EncodeToString(EmptyMd5Hash[:])
)

func Md5(str string) [16]byte {
	data := []byte(str)

	return md5.Sum(data)
}

func Md5Ptr(str *string) [16]byte {
	if str == nil {
		return EmptyMd5Hash
	}

	data := []byte(*str)

	return md5.Sum(data)
}

func Md5HashString(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Returns EmptyMd5HashString on nil any or.
//
// Any parsing error then panic.
func Md5HashAnyToString(any interface{}) string {
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	if allBytes == nil {
		return EmptyMd5HashString
	}

	hash := md5.Sum(*allBytes)
	return hex.EncodeToString(hash[:])
}
