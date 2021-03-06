// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package extime

import (
	"strings"
	"time"

	"github.com/ysicing/go-utils/convert"
)

// NowUnixString 当前时间时间戳
func NowUnixString() string {
	return convert.Int642Str(time.Now().Unix())
}

// NowUnix 当前时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// NowFormat 当前时间format
func NowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Today0hourUnix 今天0时时间戳
func Today0hourUnix() int64 {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, -1)
	return t1.Unix()
}

// BeforeNowUnix 历史时间戳
func BeforeNowUnix(old string) (oldunix int64) {
	return time.Now().Unix() - convert.Str2Int64(old)
}

// UnixInt642String unix转化为字符串
func UnixInt642String(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

// UnixString2String unix转化为字符串
func UnixString2String(t string) string {
	return time.Unix(convert.Str2Int64(t), 0).Format("2006-01-02 15:04:05")
}

// UnixNanoInt642String unix转化为字符串
func UnixNanoInt642String(t int64) string {
	return time.Unix(0, t).Format("2006-01-02 15:04:05")
}

// UnixNanoString2String unix转化为字符串
func UnixNanoString2String(t string) string {
	return time.Unix(0, convert.Str2Int64(t)).Format("2006-01-02 15:04:05")
}

// GetTodayMin 获取今天时间分钟
func GetTodayMin() string {
	return time.Now().Format("200601021504")
}

// GetTodayHour 获取今天时间小时
func GetTodayHour() string {
	return time.Now().Format("2006010215")
}

// GetToday 获取今天时间
func GetToday() string {
	return time.Now().Format("20060102")
}

// GetMonth 获取当前月份
func GetMonth() string {
	return time.Now().Format("200601")
}

// GetYear 获取当前年份
func GetYear() string {
	return time.Now().Format("2006")
}

// GetWeekFristDayUnix 时间
func GetWeekFristDayUnix() int64 {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart.Unix()
}

// GetWeekLastDayUnix 时间
func GetWeekLastDayUnix() int64 {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekNextStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
	return weekNextStart.Unix() - 1
}

// GetWeekDayUnix 时间
func GetWeekDayUnix() (int64, int64) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekNextStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
	return weekStart.Unix(), weekNextStart.Unix() - 1
}

// GetWeekDayUnixString 时间
func GetWeekDayUnixString() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekNextStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+7)
	return weekStart.Format("2006-01-02 15:03:04"), weekNextStart.Format("2006-01-02 15:03:04")
}

// NowAddUnix2Int64 当前时间戳 add
func NowAddUnix2Int64(key string, value time.Duration) int64 {
	lkey := strings.ToLower(key)
	if lkey == "m" || lkey == "minute" {
		return time.Now().Add(time.Minute * value).Unix()
	}
	if lkey == "d" || lkey == "day" {
		return time.Now().Add(time.Hour * 24 * value).Unix()
	}
	if lkey == "w" || lkey == "week" {
		return time.Now().Add(time.Hour * 24 * 7 * value).Unix()
	}
	return time.Now().Add(time.Hour * value).Unix()
}

// NowAddUnix2Str 当前时间戳 add
func NowAddUnix2Str(key string, value time.Duration) string {
	lkey := strings.ToLower(key)
	if lkey == "m" || lkey == "minute" {
		return time.Now().Add(time.Minute * value).String()
	}
	if lkey == "d" || lkey == "day" {
		return time.Now().Add(time.Hour * 24 * value).String()
	}
	if lkey == "w" || lkey == "week" {
		return time.Now().Add(time.Hour * 24 * 7 * value).String()
	}
	return time.Now().Add(time.Hour * value).String()
}
