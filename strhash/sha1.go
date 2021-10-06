package strhash

import (
	"crypto/sha1"
	"encoding/hex"

	"gitlab.com/evatix-go/strhelper/anyto"
)

var (
	// sha1 is also same for nil or "" https://play.golang.org/p/oMAxxSzAP7F
	internalSha1        = sha1.New()
	EmptySha1Hash       = internalSha1.Sum([]byte(nil))
	EmptySha1HashString = hex.EncodeToString(EmptySha1Hash[:])
)

func Sha1(str string) []byte {
	if str == "" {
		return EmptySha1Hash
	}

	bytes := []byte(str)

	return internalSha1.Sum(bytes)
}

func Sha1Ptr(str *string) []byte {
	if str == nil || *str == "" {
		return EmptySha1Hash
	}

	bytes := []byte(*str)

	return internalSha1.Sum(bytes)
}

func Sha1String(text string) string {
	hash := internalSha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Returns EmptySha1Hash on nil any or.
//
// Any parsing error then panic.
func Sha1Any(any interface{}) []byte {
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	if allBytes == nil {
		return EmptySha1Hash
	}

	return internalSha1.Sum(allBytes)
}

// Returns EmptySha1HashString on nil any or.
//
// Any parsing error then panic.
func Sha1AnyToString(any interface{}) string {
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	if allBytes == nil {
		return EmptySha1HashString
	}

	hash := internalSha1.Sum(allBytes)
	return hex.EncodeToString(hash[:])
}
