package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DateFmtYMDHIS = "2006-01-02 15:04:05"

func GetSpecialEndDate(Times time.Time, formatString string) (time.Time, error) {
	var resp time.Time
	sec, _ := time.ParseDuration("-1s")
	switch {
	case strings.HasSuffix(formatString, "d"):
		n, _ := strconv.Atoi(strings.TrimSuffix(formatString, "d"))
		currentYear, currentMonth, currentDay := Times.Date()
		currentLocation := Times.Location()
		firstOfDay := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
		resp = firstOfDay.AddDate(0, 0, n)
	case strings.HasSuffix(formatString, "m"):
		n, _ := strconv.Atoi(strings.TrimSuffix(formatString, "m"))
		currentYear, currentMonth, _ := Times.Date()
		currentLocation := Times.Location()
		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		resp = firstOfMonth.AddDate(0, n, 0)
	case strings.HasSuffix(formatString, "y"):
		n, _ := strconv.Atoi(strings.TrimSuffix(formatString, "y"))
		currentYear, _, _ := Times.Date()
		currentLocation := Times.Location()
		firstOfDay := time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentLocation)
		resp = firstOfDay.AddDate(n, 0, 0)
	default:
		return Times, fmt.Errorf("not Supported Yet")
	}
	return resp.Add(sec), nil
}

// FormatDateYMDHISToYMD 将"2006-01-02 15:04:05" 格式的时间转为"2006-01-02"
func FormatDateYMDHISToYMD(YMDHISDate string) string {
	if len(YMDHISDate) >= 10 {
		YMDData := YMDHISDate[:10]
		if YMDData == "0001-01-01" {
			return ""
		}
		return YMDData
	}
	return YMDHISDate
}

// SetNullData 过滤掉"0001-01-01"格式时间
func SetNullData(YMDHISDate string) string {
	if len(YMDHISDate) >= 10 {
		YMDData := YMDHISDate[:10]
		if YMDData == "0001-01-01" {
			return ""
		}
		return YMDHISDate
	}
	return YMDHISDate
}

// 获取月份起止时间
func GetMonthStartEnd(date string, layout string) map[string]string {
	res := make(map[string]string, 2)
	Date, _ := time.Parse(layout, date)
	res["start"] = time.Date(Date.Year(), Date.Month(), 1, 0, 0, 0, 0, time.Local).Format(DateFmtYMDHIS)
	res["end"] = time.Date(Date.Year(), Date.Month()+1, 0, 23, 59, 59, 59, time.Local).Format(DateFmtYMDHIS)
	return res
}

func GetNowYearMonth() string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	return fmt.Sprintf("%v-%v", year, month)
}

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

/* 时间格式字符串转换 */
func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}

func GetDaysBetween2Date(format, date1Str, date2Str string) (int, error) {
	// 将字符串转化为Time格式
	date1, err := time.ParseInLocation(format, date1Str, time.Local)
	if err != nil {
		return 0, err
	}
	// 将字符串转化为Time格式
	date2, err := time.ParseInLocation(format, date2Str, time.Local)
	if err != nil {
		return 0, err
	}
	//计算相差天数
	return int(date1.Sub(date2).Hours() / 24), nil
}

// 字符串转time.Time
func GetTime(timeStr string) (time.Time, error) {
	timeFormat, err := time.ParseInLocation(DateFmtYMDHIS, timeStr, time.Local)
	if err != nil {
		return time.Time{}, err
	}
	return timeFormat, nil
}

/*
GOOS=linux GOARCH=amd64 go build -o
mysql类型timestamp(时间戳)和datetime(时间)差别：

timestamp占用4个字节;
datetime占用8个字节;
timestamp范围1970-01-01 00:00:01.000000 到 2038-01-19 03:14:07.999999;
datetime是1000-01-01 00:00:00.000000 到 9999-12-31 23:59:59.999999;
时间戳格式：

10位数的时间戳是以 秒 为单位，如：1530027865
13位数的时间戳是以 毫秒 为单位， 如：1530027865231
19位数的时间戳是以 纳秒 为单位，如：1530027865231834600
Golang获取当前时间或时间戳
*/
func TimeToUnix(e time.Time) int64 {
	return e.UnixNano() / 1e6
}

func UnixToTime(e string) (datatime time.Time, err error) {
	data, err := strconv.ParseInt(e, 10, 64)
	datatime = time.Unix(data/1000, 0)
	return
}
