package utils

import (
	"strconv"
	"time"
)

//StrToInt string 转int
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

//StrToUInt string 转int
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// StrToTime 字符串转time
func StrToTime(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(timeLayout, s, loc)
	return t
}

// StrToDate 字符串转time
func StrToDate(s string) time.Time {
	loc, _ := time.LoadLocation("Local")
	timeLayout := "2006-01-02"
	t, _ := time.ParseInLocation(timeLayout, s, loc)
	return t
}

//StrToTimePtr ..
func StrToTimePtr(str string) *time.Time {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return nil
	}
	t, e := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if e != nil {
		return nil
	}
	return &t
}

