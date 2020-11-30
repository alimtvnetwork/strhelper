package strhash

import (
	"crypto/sha512"
	"encoding/hex"

	"gitlab.com/evatix-go/strhelper/anyto"
)

var (
	internalSha512        = sha512.New()
	EmptySha512Hash       = internalSha512.Sum([]byte(nil))
	EmptySha512HashString = hex.EncodeToString(EmptySha512Hash[:])
)

func Sha512(str string) []byte {
	if str == "" {
		return EmptySha512Hash
	}

	bytes := []byte(str)

	return internalSha512.Sum(bytes)
}

func Sha512Ptr(str *string) []byte {
	if str == nil || *str == "" {
		return EmptySha512Hash
	}

	bytes := []byte(*str)

	return internalSha512.Sum(bytes)
}

func Sha512String(text string) string {
	hash := internalSha512.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Returns EmptySha512Hash on nil any or.
//
// Any parsing error then panic.
func Sha512Any(any interface{}) []byte {
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	if allBytes == nil {
		return EmptySha512Hash
	}

	return internalSha512.Sum(*allBytes)
}

// Returns EmptySha512HashString on nil any or.
//
// Any parsing error then panic.
func Sha512AnyToString(any interface{}) string {
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	if allBytes == nil {
		return EmptySha512HashString
	}

	hash := internalSha512.Sum(*allBytes)
	return hex.EncodeToString(hash[:])
}
