package util

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
)

func GetBytesSha256(bytes []byte) string {
	hash := sha256.New()
	hash.Write(bytes)
	return hex.EncodeToString(hash.Sum(nil))
}

func UUIDTOID (uuid string) string {
	if len(uuid) < 36 {
		return ""
	}
	str := uuid[:36]
	var id string
	id = strings.ReplaceAll(str, "-", "")
	return id
}

func IDTOUUID (id string) string {
	if len(id) < 32 {
		return ""
	}
	str := []string{id[:8], id[8:12], id[12:16], id[16:20], id[20:]}
	uuid := strings.Join(str, "-")
	return uuid
}

func PathExist(path string) bool {
	if s, err := os.Stat(path); err != nil {
		return false
	} else {
		if s.IsDir() {
			return false
		}
		return true
	}
}
