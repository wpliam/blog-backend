package util

import (
	"blog-backend/constant"
	"time"
)

// ParseMysqlDate 解析mysql的时间为字符串
func ParseMysqlDate(layout string, value string) string {
	date, err := time.Parse(constant.MysqlDefaultTimeLayout, value)
	if err != nil {
		return value
	}
	return date.Format(layout)
}

// ParseTimeUnix 解析时间为时间戳
func ParseTimeUnix(layout string, value string) int64 {
	date, err := time.Parse(layout, value)
	if err != nil {
		return 0
	}
	return date.Unix()
}
