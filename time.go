package itools

import (
	"time"
)

// CurrentDayZero 获取当天零点时间
func CurrentDayZero() string {
	return time.Now().Format(time.DateOnly) + " 00:00:00"
}

// CurrentDayEnd 获取当天结束时间
func CurrentDayEnd() string {
	return time.Now().Format(time.DateOnly) + " 23:59:59"
}

// CalculateCurrentTimeAndZeroTime 计算当前时间和零点的时间差
func CalculateCurrentTimeAndZeroTime() time.Duration {
	now := time.Now().Unix()
	nd := time.Now().Format(time.DateOnly) + " 23:59:59"
	targetTime, _ := time.ParseInLocation(time.DateOnly, nd, time.Local)
	ts := targetTime.Unix()
	return time.Duration(ts-now) * time.Second

}

// PreviousDayStartTime 获取前一天零点时间
func PreviousDayStartTime() string {
	now := time.Now()
	old := now.AddDate(0, 0, -1).Format(time.DateOnly)
	return old + " 00:00:00"
}

// PreviousDayEndTime 获取前一天结束时间
func PreviousDayEndTime() string {
	now := time.Now()
	old := now.AddDate(0, 0, -1).Format(time.DateOnly)
	return old + " 23:59:59"
}

// PreviousAfterDate 获取前后指定的时间日期
func PreviousAfterDate(n int) string {
	now := time.Now()
	t := now.AddDate(0, 0, n).Format(time.DateOnly)
	return t
}

// PreviousAfterTime 获取前后指定的日期的当前时间
func PreviousAfterTime(n int) string {
	now := time.Now()
	t := now.AddDate(0, 0, n).Format(time.DateOnly)
	return t
}
