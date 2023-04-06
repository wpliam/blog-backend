package util

import (
	"github.com/wpliap/common-wrap/log"
	"strconv"
	"strings"
)

func ParseInt64(str string) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		log.Errorf("ParseInt64 err:%v str:%s", err, str)
		return 0
	}
	return val
}

func ParseArrInt64(str string) []int64 {
	ids := make([]int64, 0)
	arr := strings.Split(str, ",")
	for _, s := range arr {
		id := ParseInt64(s)
		if id == 0 {
			continue
		}
		ids = append(ids, id)
	}
	return ids
}
