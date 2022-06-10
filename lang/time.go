package lang

import "time"

func GetNow() *time.Time {
	now := time.Now()
	return &now
}
