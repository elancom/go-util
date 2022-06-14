package lang

import "time"

func GetNow() *time.Time {
	now := time.Now()
	return &now
}

const formatDate = "2006-01-02"
const formatDateTime = "2006-01-02 15:04:05"
const formatTime = "15:04:05"

func FormatDateTimes(t time.Time) string {
	return FormatDateTime(&t)
}

func FormatDateTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(formatDateTime)
}

func FormatDates(t time.Time) string {
	return FormatDateTime(&t)
}

func FormatDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(formatDate)
}

func FormatTimes(t time.Time) string {
	return FormatTime(&t)
}

func FormatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(formatTime)
}
