package stringid

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"time"
)

var (
	validHex = regexp.MustCompile(`^[a-f0-9]{64}$`)
)

func init() {
	// 初始化全局随机数种子
	var seed int64
	buffer := make([]byte, 8)
	if _, err := cryptorand.Read(buffer); err != nil {
		seed = time.Now().UnixNano()
	} else {
		seed = int64(binary.BigEndian.Uint64(buffer))
	}

	rand.Seed(seed)
}

// 生成256bit随机ID
func GenerateRandomID() string {
	buffer := make([]byte, 32)
	if _, err := cryptorand.Read(buffer); err != nil {
		rand.Read(buffer)
	}
	for {
		id, success := toHexID(buffer)
		if success {
			return id
		}
	}
}

// 生成256bit伪随机ID
func GenerateNonCryptoID() string {
	buffer := make([]byte, 32)
	rand.Read(buffer)
	for {
		id, success := toHexID(buffer)
		if success {
			return id
		}
	}
}

func toHexID(buffer []byte) (string, bool) {
	id := hex.EncodeToString(buffer)
	// 如果要将id作为hostname，必须判断是否全为数字，某些语言处理全数字hostname时会出错
	// https://github.com/moby/moby/issues/3869
	//if _, err := strconv.ParseInt(TruncateID(id), 10, 64); err == nil {
	//	return "", false
	//}
	return id, true
}

// 验证ID是否有效
func ValidateID(id string) error {
	if ok := validHex.MatchString(id); !ok {
		return fmt.Errorf("ID %q is invalid", id)
	}
	return nil
}

func GetStanderUUID() (string, error) {
	data, err := ioutil.ReadFile("/proc/sys/kernel/random/uuid")
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
