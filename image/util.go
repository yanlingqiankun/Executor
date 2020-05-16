package image

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

//得到一个文件的哈希值
func getSha256(filePath string) (string, error) {
	var hashValue string
	file, err := os.Open(filePath)
	if err != nil {
		return hashValue, err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashValue, err
	}
	hashInBytes := hash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes)
	return hashValue, nil
}

func returnWithError(info string, err error) {
	if err != nil {
		logger.WithError(err).Error(info)
	}
}

func exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}