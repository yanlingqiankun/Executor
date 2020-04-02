package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetBytesSha256(bytes []byte) string {
	hash := sha256.New()
	hash.Write(bytes)
	return hex.EncodeToString(hash.Sum(nil))
}
