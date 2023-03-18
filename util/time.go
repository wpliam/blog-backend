package util

import (
	"blog-backend/constant"
	"time"
)

func ParseDateTime(layout string, value string) string {
	date, err := time.Parse(constant.MysqlDefaultTimeLayout, value)
	if err != nil {
		return value
	}
	return date.Format(layout)
}
