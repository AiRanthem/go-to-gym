package gym

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

type TimeFunc func() (string, string)

var WeekDayMap = map[string]string{
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五",
	"Saturday":  "星期六",
	"Sunday":    "星期天",
}

func Lunch() (string, string) {
	return time.Now().Format("2006-01-02,") + WeekDayMap[time.Now().Format("Monday")], "12:00-13:30"
}

func Nap() (string, string) {
	return time.Now().Format("2006-01-02,") + WeekDayMap[time.Now().Format("Monday")], "13:30-15:00"
}

func Tea() (string, string) {
	return time.Now().Format("2006-01-02,") + WeekDayMap[time.Now().Format("Monday")], "15:00-16:30"
}

func Work() (string, string) {
	return time.Now().Format("2006-01-02,") + WeekDayMap[time.Now().Format("Monday")], "16:30-18:00"
}

func Dinner() (string, string) {
	return time.Now().Format("2006-01-02,") + WeekDayMap[time.Now().Format("Monday")], "18:00-19:30"
}

func Night() (string, string) {
	return time.Now().Format("2006-01-02,") + WeekDayMap[time.Now().Format("Monday")], "19:30-21:00"
}

func Now() (string, string) {
	h, _ := strconv.Atoi(time.Now().Format("15"))
	m, _ := strconv.Atoi(time.Now().Format("4"))
	log.WithField("hour", h).WithField("minute", m).Info("You are so FUSSY! Hurry UP and GO")
	t := h*60 + m
	if t >= 19*60+30 {
		return Night()
	}
	if t >= 18*60 {
		return Dinner()
	}
	if t >= 16*60+30 {
		return Work()
	}
	if t >= 15*60 {
		return Tea()
	}
	if t >= 13*60+30 {
		return Nap()
	}
	return Lunch()
}

func GetTimeFunc(name string) (TimeFunc, error) {
	switch strings.ToLower(name) {
	case "lunch":
		return Lunch, nil
	case "nap":
		return Nap, nil
	case "tea":
		return Tea, nil
	case "work":
		return Work, nil
	case "dinner":
		return Dinner, nil
	case "night":
		return Night, nil
	case "now":
		return Now, nil
	default:
		return nil, fmt.Errorf("no such time func %s", name)
	}
}
