package util

import (
	"github.com/wpliap/common-wrap/log"
	"strconv"
)

func ParseInt64(str string) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Errorf("ParseInt64 err:%v str:%s", err, str)
		return 0
	}
	return val
}
