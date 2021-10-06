package ktime

import (
	"time"
)

// GetTodayTimestamp 返回 YYYYMMDD
func GetTodayTimestamp() (int64, string) {
	now := time.Now()
	loc, _ := time.LoadLocation(LocalTimeZone)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc).Unix()
	todayStr := now.Format(DateSimp)
	return today, todayStr
}

// GetYesterdayTimestamp 返回 YYYYMMDD
func GetYesterdayTimestamp() (int64, string) {
	now := time.Now()
	loc, _ := time.LoadLocation(LocalTimeZone)
	yesterday := now.AddDate(0, 0, -1)
	yesterdayBegin := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, loc).Unix()
	yesterdayStr := yesterday.Format(DateSimp)
	return yesterdayBegin, yesterdayStr
}

// GetToday 返回今天 0点的时间戳
func GetToday() time.Time {
	return GetDayBefore(0)
}

// GetYestorday 返回昨天0点的时间戳
func GetYestorday() time.Time {
	return GetDayBefore(1)
}

// GetDayBefore 返回n天之前
func GetDayBefore(n int) time.Time {
	day := time.Now().Add(time.Duration(n) * time.Hour * -24)
	loc, _ := time.LoadLocation(LocalTimeZone)
	return time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
}

// GetDayAfter 返回n天之后
func GetDayAfter(n int) time.Time {
	return GetDayBefore(-n)
}

// GetTodayLastSecTime 返回当天最后一秒的时间
func GetTodayLastSecTime() time.Time {
	now := time.Now()
	loc, _ := time.LoadLocation(LocalTimeZone)
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, loc)
}

// ToMySQLFormat 返回mysql里存的时间格式
func ToMySQLFormat(t time.Time) string {
	return t.Format(DateTimeHyphen)
}

// ParseMySQLFormat 按照mysql里存的时间格式 parse
func ParseMySQLFormat(s string) (time.Time, error) {
	loc, _ := time.LoadLocation(LocalTimeZone)
	return time.ParseInLocation(DateTimeHyphen, s, loc)
}

// NextMondayMorning 返回下一周周一00:01的时间戳
func NextMondayMorning() time.Time {
	today := GetToday()
	todayDay := today.Weekday()
	// 算到周六差几天
	durationDay := int(time.Saturday) - int(todayDay)
	// + 2 day 就是到周一 00:00
	monday0001 := today.Add(time.Duration((durationDay+2)*24) * time.Hour).Add(time.Minute)
	return monday0001

	// another way...
	// seconds := time.Now().Unix()
	// var secondsOfWeek int64 = (60 * 60 * 24 * 7)
	// weeks := seconds / secondsOfWeek

	// sunday2400_2 := time.Unix((weeks)*secondsOfWeek, 0).Add(4 * 24 * time.Hour).Add(-8 * time.Hour)
	// fmt.Printf("%v\n", sunday2400_2)
}

// NextMonday8AM 下周一早8点
func NextMonday8AM() time.Time {
	return NextMondayMorning().Add(-1 * time.Minute).Add(8 * time.Hour)
}

// ThisMondayMorning 返回本周周一00:00的时间戳
func ThisMondayMorning() time.Time {
	return NextMondayMorning().Add(-7 * 24 * time.Hour)
}
