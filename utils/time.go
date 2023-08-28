package utils

import (
	"airport/defines"
	"fmt"
	"math"
	"time"
)

func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(defines.TimeFormat)
}

func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(defines.DateFormat)
}

func ParseDateTime(str string, loc ...*time.Location) (t time.Time, err error) {
	base := "0000-00-00 00:00:00.0000000"
	timeFormat := "2006-01-02 15:04:05.999999"

	switch len(str) {
	case 10, 19, 21, 22, 23, 24, 25, 26: // up to "YYYY-MM-DD HH:MM:SS.MMMMMM"
		if str == base[:len(str)] {
			return
		}
		t, err = time.Parse(timeFormat[:len(str)], str)
	default:
		err = fmt.Errorf("invalid time string: %s", str)
		return
	}

	// Adjust location
	if err == nil && len(loc) > 0 && loc[0] != time.UTC {
		y, mo, d := t.Date()
		h, mi, s := t.Clock()
		t, err = time.Date(y, mo, d, h, mi, s, t.Nanosecond(), loc[0]), nil
	}

	return
}

func TimeToInt64(text string) int64 {
	time, err := ParseDateTime(text)
	if err != nil {
		return 0
	}

	return time.Unix()
}

func Int64ToTimeString(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}

	return FormatTime(time.Unix(timestamp, 0))
}

func Int64ToDateString(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}

	return FormatDate(time.Unix(timestamp, 0))
}

func GetTime(input int64) time.Time {
	maxd := time.Duration(math.MaxInt64).Truncate(100 * time.Nanosecond)
	maxdUnits := int64(maxd / 100) // number of 100-ns units

	if input == 0 {
		return time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	t := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	for input > maxdUnits {
		t = t.Add(maxd)
		input -= maxdUnits
	}
	if input != 0 {
		t = t.Add(time.Duration(input * 100))
	}
	return t
}
