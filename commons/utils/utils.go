package utils

import (
	"strconv"
	"time"
)

const timezone = "Asia/Kolkata"

func GetCurrentTimestampInMillis() int64 {
	loc, _ := time.LoadLocation(timezone)
	return time.Now().In(loc).UnixMilli()
}

func GetTimeTakenInString(start time.Time) string {
	return strconv.Itoa(int(time.Since(start).Milliseconds()))
}