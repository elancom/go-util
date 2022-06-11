package lang

import "time"

func GetNow() *time.Time {
	now := time.Now()
	return &now
}

const formatDate = "2006-01-02"
const formatDateTime = "2006-01-02 15:04:05"
const formatTime = "15:04:05"

func FormatDateTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(formatDateTime)
}

func FormatDate(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(formatDate)
}

func FormatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(formatTime)
}
