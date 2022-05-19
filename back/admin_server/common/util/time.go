package util

import (
	"goods/common/constant"
	"strconv"
	"time"
)

// 获取时间
func NowTime() time.Time {
	return time.Now()
}

// 获取本地时区
func LocalTimeZone() *time.Location {
	return time.Local
}

// 使用本地时区解析时间
func ParseTime(layout, datetime string) (time.Time, error) {
	return time.ParseInLocation(layout, datetime, time.Local)
}

// 获取当前日期时间
func Now() string {
	return GetDateTime(0, constant.DateTimeLayout)
}

// 获取日期时间
func GetDateTime(add int64, layout string) string {
	if add == 0 {
		return NowTime().Format(layout)
	}
	t := time.Unix(GetTimestamp(add), 0)
	return t.Format(layout)
}

// 获取时间戳
func GetTimestamp(add int64) int64 {
	return NowTime().Unix() + add
}

// 根据时间戳获取日期时间
func GetDateTimeByTimestamp(timestamp int64, layout string) string {
	t := time.Unix(timestamp, 0)
	return t.Format(layout)
}

// 根据日期时间获取时间戳
func GetTimeStampByDateTime(datetime string, layout string) int64 {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation(layout, datetime, loc)
	return t.Unix()
}

// 根据日期格式化日期。
// 例如:
//   datetime的格式是YYYY-mm-dd 00:00:00  则selfLayout传 "2006-01-02 15:04:05"，两者的格式要保持一致
//   YYYY-mm-dd 则selfLayout传 "2006-01-02"
//   changeLayout 则是你需要更改的格式
func FormatStringTimeByDataTime(datetime string, selfLayout string, changeLayout string) string {
	tmpTime := GetTimeStampByDateTime(datetime, selfLayout)
	return GetDateTimeByTimestamp(tmpTime, changeLayout)
}

// 获取明天第一秒的time
func TomorrowFirstTime(now ...time.Time) time.Time {
	nowTime := time.Now()
	if len(now) > 0 {
		nowTime = now[0]
	}
	tomorrowFirst, _ := ParseTime(constant.DateTimeLayout, nowTime.AddDate(0, 0, 1).Format(constant.DateLayout)+" 00:00:00")
	return tomorrowFirst
}

// 获取明天第一秒的timestamp
func TomorrowFirstTimestamp(now ...time.Time) int64 {
	return TomorrowFirstTime(now...).Unix()
}

// 获取今天最后一秒的time
func TodayLastTime(now ...time.Time) time.Time {
	nowTime := time.Now()
	if len(now) > 0 {
		nowTime = now[0]
	}
	todayLast, _ := ParseTime(constant.DateTimeLayout, nowTime.Format(constant.DateLayout)+" 23:59:59")
	return todayLast
}

// 获取今天最后一秒的timestamp
func TodayLastTimestamp(now ...time.Time) int64 {
	return TodayLastTime(now...).Unix()
}

// 获取当天剩余秒数
func TodayRestSecond(now ...time.Time) int64 {
	nowTime := time.Now()
	if len(now) > 0 {
		nowTime = now[0]
	}
	nowUnix := nowTime.Unix()
	tomorrowFirstUnix := TomorrowFirstTimestamp(nowTime)
	return tomorrowFirstUnix - nowUnix
}

// 根据年份获取年龄
func GetAgeByYear(year string) (age int) {
	formatYear, _ := strconv.Atoi(year)
	if formatYear <= 0 {
		age = -1
	}
	nowYear := time.Now().Year()
	age = nowYear - formatYear
	return
}
