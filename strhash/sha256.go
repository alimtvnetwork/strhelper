package strhash

import (
	"crypto/sha256"
	"encoding/hex"

	"gitlab.com/evatix-go/strhelper/anyto"
)

var (
	internalSha256        = sha256.New()
	EmptySha256Hash       = internalSha256.Sum([]byte(nil))
	EmptySha256HashString = hex.EncodeToString(EmptySha256Hash[:])
)

func Sha256(str string) []byte {
	if str == "" {
		return EmptySha256Hash
	}

	bytes := []byte(str)

	return internalSha256.Sum(bytes)
}

func Sha256Ptr(str *string) []byte {
	if str == nil || *str == "" {
		return EmptySha256Hash
	}

	bytes := []byte(*str)

	return internalSha256.Sum(bytes)
}

func Sha256String(text string) string {
	hash := internalSha256.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Returns EmptySha256Hash on nil any or.
//
// Any parsing error then panic.
func Sha256Any(any interface{}) []byte {
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	if allBytes == nil {
		return EmptySha256Hash
	}

	return internalSha256.Sum(*allBytes)
}

// Returns EmptySha256HashString on nil any or.
//
// Any parsing error then panic.
func Sha256AnyToString(any interface{}) string {
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	if allBytes == nil {
		return EmptySha256HashString
	}

	hash := internalSha256.Sum(*allBytes)
	return hex.EncodeToString(hash[:])
}
