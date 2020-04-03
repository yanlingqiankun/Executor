package utils

import (
	"fmt"
	"github.com/yanlingqiankun/Executor/pb"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

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

// 如果存在错误则打印错误信息并返回true，否则直接返回false
func PrintError(err *pb.Error) bool {
	if err == nil {
		return false
	} else {
		if err.Code == 0 {
			return false
		}
		fmt.Printf("[ERR %d] %s", err.Code, err.Message)
		return true
	}
}

func CheckLength(param string) string {
	if len(param) >= 12 {
		return param[:12]
	}
	return param
}

func getWindowSize() [2]int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	strSize := string(out)
	strSize = strings.TrimSuffix(strSize, "\n")
	strListSize := strings.Split(strSize, " ")
	sizes := [2]int{}
	for index, str := range strListSize {
		if size_one, err := strconv.Atoi(str); err != nil {
			panic(err)
		} else {
			sizes[index] = size_one
		}
	}
	return sizes
}